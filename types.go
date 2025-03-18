package go_limitless

import "time"

type Lifelog struct {
	ID       string        `json:"id"`
	Title    string        `json:"title"`
	Markdown string        `json:"markdown,omitempty"`
	Contents []ContentNode `json:"contents"`
}

type ContentNode struct {
	Type              string        `json:"type"`
	Content           string        `json:"content"`
	StartTime         string        `json:"startTime"`
	EndTime           string        `json:"endTime"`
	StartOffsetMs     int           `json:"startOffsetMs"`
	EndOffsetMs       int           `json:"endOffsetMs"`
	Children          []ContentNode `json:"children,omitempty"`
	SpeakerName       *string       `json:"speakerName,omitempty"`
	SpeakerIdentifier *string       `json:"speakerIdentifier,omitempty"`
}

type MetaLifelogs struct {
	NextCursor *string `json:"nextCursor,omitempty"`
	Count      int     `json:"count"`
}

type LifelogsResponse struct {
	Data struct {
		Lifelogs []Lifelog `json:"lifelogs"`
	} `json:"data"`
	Meta struct {
		Lifelogs MetaLifelogs `json:"lifelogs"`
	} `json:"meta"`
}

type GetLifelogsParams struct {
	Timezone        string    `url:"timezone,omitempty"`        // IANA timezone specifier
	Date            string    `url:"date,omitempty"`            // YYYY-MM-DD format
	Start           time.Time `url:"start,omitempty"`           // Start of time range
	End             time.Time `url:"end,omitempty"`             // End of time range
	Cursor          string    `url:"cursor,omitempty"`          // Pagination cursor
	Direction       string    `url:"direction,omitempty"`       // asc or desc
	IncludeMarkdown *bool     `url:"includeMarkdown,omitempty"` // Include markdown content
	IncludeHeadings *bool     `url:"includeHeadings,omitempty"` // Include headings
	Limit           int       `url:"limit,omitempty"`           // Max results to return
}
