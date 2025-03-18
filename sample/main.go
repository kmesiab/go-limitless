package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	limitless "github.com/kmesiab/go-limitless"
)

func main() {
	apiKey := os.Getenv("LIMITLESS_API_KEY")
	if apiKey == "" {
		log.Fatal("LIMITLESS_API_KEY environment variable is required")
	}

	// Create client with API key
	client := limitless.NewClient(apiKey)

	// Define query parameters using the proper struct
	params := &limitless.GetLifelogsParams{
		Timezone:        "America/New_York",
		Date:            "2023-12-25",                                  // Single day
		Start:           time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC), // Or date range
		End:             time.Date(2023, 12, 26, 0, 0, 0, 0, time.UTC),
		Direction:       "desc",
		IncludeMarkdown: limitless.BoolPtr(true),
		IncludeHeadings: limitless.BoolPtr(true),
		Limit:           50,
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
