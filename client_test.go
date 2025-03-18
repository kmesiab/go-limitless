package go_limitless

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetLifelogs_Success(t *testing.T) {
	mockResponse := `{
        "data": {
            "lifelogs": [{ "id": "123", "title": "Test Entry", "markdown": "# Heading", "contents": [] }]
        },
        "meta": {
            "lifelogs": { "count": 1 }
        }
    }`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(mockResponse))
		if err != nil {
			return
		}
	}))
	defer ts.Close()

	client := NewClient("test-api-key")
	client.BaseURL = ts.URL

	ctx := context.Background()
	params := &GetLifelogsParams{
		Limit: 1,
	}
	resp, err := client.GetLifelogs(ctx, params)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(resp.Data.Lifelogs) != 1 {
		t.Fatalf("expected 1 lifelog, got %d", len(resp.Data.Lifelogs))
	}
}

func TestGetLifelogs_ErrorResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer ts.Close()

	client := NewClient("invalid-api-key")
	client.BaseURL = ts.URL

	ctx := context.Background()
	params := &GetLifelogsParams{
		Limit: 1,
	}
	_, err := client.GetLifelogs(ctx, params)

	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetLifelogs_InvalidJSON(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("invalid json"))
		if err != nil {
			return
		}
	}))
	defer ts.Close()

	client := NewClient("test-api-key")
	client.BaseURL = ts.URL

	ctx := context.Background()
	params := &GetLifelogsParams{
		Limit: 1,
	}
	_, err := client.GetLifelogs(ctx, params)

	if err == nil {
		t.Fatal("expected a JSON decoding error but got none")
	}
}

func TestGetLifelogs_Timeout(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	client := NewClient("test-api-key")
	client.HTTPClient.Timeout = 1 * time.Second // Force timeout
	client.BaseURL = ts.URL

	ctx := context.Background()
	params := &GetLifelogsParams{
		Limit: 1,
	}
	_, err := client.GetLifelogs(ctx, params)

	if err == nil {
		t.Fatal("expected timeout error but got none")
	}
}

func TestGetLifelogs_MissingAuthToken(t *testing.T) {
	client := NewClient("") // Empty API Key
	ctx := context.Background()
	params := &GetLifelogsParams{
		Limit: 1,
	}
	_, err := client.GetLifelogs(ctx, params)

	if err == nil {
		t.Fatal("expected authentication error but got none")
	}
}

func TestGetLifelogs_BadRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer ts.Close()

	client := NewClient("test-api-key")
	client.BaseURL = ts.URL

	ctx := context.Background()
	params := &GetLifelogsParams{
		Limit: -1, // Invalid limit to trigger bad request
	}
	_, err := client.GetLifelogs(ctx, params)

	if err == nil {
		t.Fatal("expected bad request error but got none")
	}
}

func TestGetLifelog_Success(t *testing.T) {
	mockResponse := `{
		"id": "test-id",
		"title": "Test Lifelog",
		"markdown": "# Test Lifelog",
		"contents": [{ "type": "heading1", "content": "Test Content" }]
	}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(mockResponse))
		if err != nil {
			return
		}
	}))
	defer ts.Close()

	client := NewClient("test-api-key")
	client.BaseURL = ts.URL

	ctx := context.Background()
	lifelog, err := client.GetLifelog(ctx, "test-id")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if lifelog.ID != "test-id" {
		t.Fatalf("expected lifelog ID 'test-id', got '%s'", lifelog.ID)
	}

	if lifelog.Title != "Test Lifelog" {
		t.Fatalf("expected title 'Test Lifelog', got '%s'", lifelog.Title)
	}
}

func TestGetLifelog_NotFound(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte("{\"error\": \"Lifelog not found\"}"))
		if err != nil {
			return
		}
	}))
	defer ts.Close()

	client := NewClient("test-api-key")
	client.BaseURL = ts.URL

	ctx := context.Background()
	_, err := client.GetLifelog(ctx, "nonexistent-id")

	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetLifelog_InvalidJSON(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("invalid json"))
		if err != nil {
			return
		}
	}))
	defer ts.Close()

	client := NewClient("test-api-key")
	client.BaseURL = ts.URL

	ctx := context.Background()
	_, err := client.GetLifelog(ctx, "test-id")

	if err == nil {
		t.Fatal("expected a JSON decoding error but got none")
	}
}

func TestGetLifelog_Timeout(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	client := NewClient("test-api-key")
	client.HTTPClient.Timeout = 1 * time.Second // Force timeout
	client.BaseURL = ts.URL

	ctx := context.Background()
	_, err := client.GetLifelog(ctx, "test-id")

	if err == nil {
		t.Fatal("expected timeout error but got none")
	}
}

func TestGetLifelog_Unauthorized(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		_, err := w.Write([]byte("{\"error\": \"Unauthorized\"}"))
		if err != nil {
			return
		}
	}))
	defer ts.Close()

	client := NewClient("invalid-api-key")
	client.BaseURL = ts.URL

	ctx := context.Background()
	_, err := client.GetLifelog(ctx, "test-id")

	if err == nil {
		t.Fatal("expected unauthorized error but got none")
	}
}

func TestGetLifelogs_Parameters(t *testing.T) {
	tests := []struct {
		name          string
		params        *GetLifelogsParams
		expectedQuery string
		mockResponse  string
		expectedError bool
	}{
		{
			name: "all parameters",
			params: &GetLifelogsParams{
				Timezone:        "America/New_York",
				Date:            "2023-12-25",
				Start:           time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC),
				End:             time.Date(2023, 12, 26, 0, 0, 0, 0, time.UTC),
				Cursor:          "next-page",
				Direction:       "desc",
				IncludeMarkdown: boolPtr(true),
				IncludeHeadings: boolPtr(false),
				Limit:           10,
			},
			expectedQuery: "cursor=next-page&date=2023-12-25&direction=desc&end=2023-12-26T00%3A00%3A00Z&includeHeadings=false&includeMarkdown=true&limit=10&start=2023-12-25T00%3A00%3A00Z&timezone=America%2FNew_York",
		},
		{
			name: "minimal parameters",
			params: &GetLifelogsParams{
				Limit: 1,
			},
			expectedQuery: "limit=1",
		},
		{
			name:          "nil parameters",
			params:        nil,
			expectedQuery: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if got := r.URL.RawQuery; got != tt.expectedQuery {
					t.Errorf("expected query %q, got %q", tt.expectedQuery, got)
				}
				w.WriteHeader(http.StatusOK)
				_, err := w.Write([]byte(`{"data":{"lifelogs":[]},"meta":{"lifelogs":{"count":0}}}`))
				if err != nil {
					return
				}
			}))
			defer ts.Close()

			client := NewClient("test-api-key")
			client.BaseURL = ts.URL

			_, err := client.GetLifelogs(context.Background(), tt.params)
			if (err != nil) != tt.expectedError {
				t.Errorf("expected error %v, got %v", tt.expectedError, err)
			}
		})
	}
}

func boolPtr(b bool) *bool {
	return &b
}
