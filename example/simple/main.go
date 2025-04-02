package main

import (
	"fmt"

	"github.com/FFengIll/typechat.go"
)

// VenueData represents a venue and its description
type VenueData struct {
	Venue       string `json:"venue"`
	Description string `json:"description"`
}

// Response represents the response structure containing a list of venues
type Response struct {
	Data []VenueData `json:"data"`
}

func main() {
	prompt := "Provide 3 suggestions for specific places to go to in Seattle on a rainy day."

	// simple api
	res, err := typechat.Traslate(prompt, &Response{}, &VenueData{})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("%+v\n", res)

	// general method
	t := typechat.NewTranslator()
	str, err := t.Generate(prompt, &Response{}, &VenueData{})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("%+v\n", str)
}
