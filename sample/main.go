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

	// Create client with API key
	client := limitless.NewClient(apiKey)

	// Define query parameters
	params := map[string]string{
		"limit": "10",
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
