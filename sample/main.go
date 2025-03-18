package main

import (
	"context"
	"fmt"
	"log"
	"os"

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
		Limit: 10,
		// Can also use other available parameters:
		// Timezone:        "America/New_York",
		// Date:            "2023-12-25",
		// Direction:       "desc",
		// IncludeMarkdown: limitless.BoolPtr(true),
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
