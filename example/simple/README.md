# Simple

# Data
## Prompt
`Provide 3 suggestions for specific places to go to in Seattle on a rainy day.`

## Generated Prompt
Provide 3 suggestions for specific places to go to in Seattle on a rainy day.
Respond strictly with JSON. The JSON should be compatible with the Go struct Response from the following:
```go
type Response struct {
        Data []main.VenueData `json:"data"`
}
type VenueData struct {
        Venue string `json:"venue"`
        Description string `json:"description"`
}
```

# Result
## GPT-3.5
```json
{
  "data": [
    {
      "venue": "Seattle Art Museum",
      "description": "Explore the extensive collection of art from around the world at the Seattle Art Museum. From contemporary art to ancient artifacts, there is something for everyone to enjoy."
    },
    {
      "venue": "Pike Place Market",
      "description": "Indulge in a unique shopping experience at Pike Place Market. Browse through local produce, crafts, and specialty shops, and enjoy a variety of delicious food options."
    },
    {
      "venue": "Chihuly Garden and Glass",
      "description": "Marvel at the stunning glass artworks created by Dale Chihuly at the Chihuly Garden and Glass exhibit. The vibrant colors and intricate designs are sure to captivate your senses."
    }
  ]
}
```
