package go_limitless

import "time"

// Lifelog A single lifelog entry
type Lifelog struct {
	// Unique identifier
	ID string

	// Title (equal to the first heading1 node)
	Title string

	// Raw markdown content (optional)
	Markdown string

	// Structured content nodes
	Contents []ContentNode
}

// ContentNode representing a section of the lifelog
type ContentNode struct {
	// Node type (heading1, heading2, heading3, blockquote, etc)
	Type string

	// Content text
	Content string

	// Timing information in ISO format for the given timezone
	StartTime string
	EndTime   string

	// Timing offsets in milliseconds from entry start
	StartOffsetMs int
	EndOffsetMs   int

	// Nested content nodes
	Children []ContentNode

	// Speaker information for certain node types (e.g., blockquote)
	SpeakerName       *string // Optional speaker identifier
	SpeakerIdentifier *string // Optional, "user" when speaker is identified as the user
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
