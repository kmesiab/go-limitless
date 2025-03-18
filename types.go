package go_limitless

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
