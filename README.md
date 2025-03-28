# Go-Limitless

![Golang](https://img.shields.io/badge/Go-00add8.svg?labelColor=171e21&style=for-the-badge&logo=go)  
![Build](https://github.com/kmesiab/go-limitless/actions/workflows/go-build.yml/badge.svg)
![Lint](https://github.com/kmesiab/go-limitless/actions/workflows/go-lint.yml/badge.svg)
![Tests](https://github.com/kmesiab/go-limitless/actions/workflows/go-test.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/kmesiab/go-limitless)](https://goreportcard.com/report/github.com/kmesiab/go-limitless)

## Limitless API Client for Go

`go-limitless` is a lightweight Go client library for interacting with the
**[Limitless API](https://www.limitless.ai/developers)**, providing access to
**lifelogs** with robust features like pagination, authentication, and error
handling.

---

## Features

✅ **Context-aware requests** for timeouts and cancellations  
✅ **Detailed error handling** with structured responses  
✅ **Pagination support** for large datasets  
✅ **Custom HTTP client support** for flexibility  
✅ **Comprehensive test coverage**

---

## Installation

To install the package, use:

```bash
go get github.com/kmesiab/go-limitless
```

---

## Authentication

You must have a **valid API key** from Limitless AI to use this client.  
Pass your API key when initializing the client:

```go
client := limitless.NewClient("your-api-key")
```

---

## Usage

### Initialize the Client

```go
package main

import (
   "context"
   "fmt"
   "log"

   limitless "github.com/kmesiab/go-limitless"
)

func main() {
   // Create client with API key
   client := limitless.NewClient("your-api-key")

   // Define query parameters
   params := &limitless.GetLifelogsParams{
      Limit: 10,
      Timezone: "America/New_York",
      Direction: "desc",
      IncludeMarkdown: limitless.BoolPtr(true),
   }

   // Fetch lifelogs
   lifelogs, err := client.GetLifelogs(context.Background(), params)
   if err != nil {
      log.Fatalf("Error fetching lifelogs: %v", err)
   }

   // Process response
   for _, log := range lifelogs.Data.Lifelogs {
      fmt.Printf("Lifelog: %s - %s\n", log.ID, log.Title)
   }
}
```

### Available Parameters

```go
type GetLifelogsParams struct {
   // IANA timezone specifier (e.g., "America/New_York")
   // If not specified, UTC is used
   Timezone string
   
   // Date in YYYY-MM-DD format
   // Returns all entries beginning on this date in the given timezone
   Date string
   
   // Start of time range in ISO-8601 format
   // YYYY-MM-DD or YYYY-MM-DD HH:mm:SS
   // Timezone/offsets are ignored
   Start time.Time
   
   // End of time range in ISO-8601 format
   // YYYY-MM-DD or YYYY-MM-DD HH:mm:SS
   // Timezone/offsets are ignored
   End time.Time
   
   // Pagination cursor for retrieving next set of entries
   Cursor string
   
   // Sort direction: "asc" or "desc" (default: "desc")
   Direction string
   
   // Whether to include markdown content (default: true)
   IncludeMarkdown *bool
   
   // Whether to include headings (default: true)
   IncludeHeadings *bool
   
   // Maximum number of entries to return
   Limit int
}
```

---

## Fetch a Single Lifelog

```go
ctx := context.Background()
lifelog, err := client.GetLifelog(ctx, "lifelog-id")
if err != nil {
     log.Fatalf("Error fetching lifelog: %v", err)
}

fmt.Printf("Lifelog ID: %s, Title: %s\n", lifelog.ID, lifelog.Title)
```

---

## Pagination Support

```go
ctx := context.Background()
params := &limitless.GetLifelogsParams{
   Limit: 10,
   Timezone: "America/New_York",
   Direction: "desc",
   IncludeMarkdown: limitless.BoolPtr(true),
}

// Fetch first page
firstPage, err := client.GetLifelogs(ctx, params)
if err != nil {
     log.Fatal(err)
}

// Fetch next page if available
if firstPage.Meta.Lifelogs.NextCursor != nil {
     params["cursor"] = *firstPage.Meta.Lifelogs.NextCursor
     nextPage, err := client.GetLifelogs(ctx, params)
     if err != nil {
          log.Fatal(err)
     }
     fmt.Println("Fetched next page of lifelogs:", len(nextPage.Data.Lifelogs))
}
```

---

## Error Handling

```go
lifelogs, err := client.GetLifelogs(ctx, params)
if err != nil {
     if apiErr, ok := err.(*limitless.ErrorResponse); ok {
          fmt.Printf("API Error: %d - %s\n", apiErr.StatusCode, apiErr.Message)
     } else {
          fmt.Printf("Request Error: %v\n", err)
     }
     return
}
```

---

## Running Tests

The package includes comprehensive tests for error handling, authentication,
and API responses.  

Run tests using:

```bash
go test ./...
```

---

## Contributing

We welcome contributions! If you'd like to help improve `go-limitless`, follow
these steps:

1. **Fork the repository** on GitHub.
2. **Clone your fork** and create a new feature branch:

   ```bash
   git clone https://github.com/your-username/go-limitless.git
   cd go-limitless
   git checkout -b feature-branch
   ```

3. **Make your changes and write tests**.
4. **Run tests locally** to ensure everything works:

   ```bash
   go test ./...
   ```

5. **Push your changes** and open a pull request:

   ```bash
   git push origin feature-branch
   ```

6. **Discuss your changes** in the PR and make any necessary revisions.

We appreciate your help in making `go-limitless` better! 🚀

---

## License

This library is distributed under the **MIT license**.
