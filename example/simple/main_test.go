package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	jsonStr := `{
        "Data": [
            {
                "venue": "Seattle Art Museum",
                "description": "Explore a vast collection of artworks spanning different cultures and time periods, perfect for immersing yourself in art on a rainy day."
            },
            {
                "venue": "Museum of Pop Culture",
                "description": "A vibrant museum dedicated to contemporary popular culture, featuring exhibits on music, science fiction, and more, ideal for entertainment indoors."
            },
            {
                "venue": "The Seattle Public Library - Central Library",
                "description": "An architectural marvel with a vast collection of books and quiet spaces, great for reading or studying while staying dry."
            }
        ]
    }`

	var response Response
	err := json.Unmarshal([]byte(jsonStr), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
		return
	}

	// Verify the data was correctly unmarshaled
	if len(response.Data) != 3 {
		t.Fatalf("Expected 3 venues, got %d", len(response.Data))
	}

	// Check first venue
	if response.Data[0].Venue != "Seattle Art Museum" {
		t.Fatalf("Expected first venue to be 'Seattle Art Museum', got '%s'", response.Data[0].Venue)
	}

	fmt.Printf("%v\n", response)
}
