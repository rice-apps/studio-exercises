package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Film represents the data structure of a film returned by the Ghibli API.
// We use struct tags (e.g. `json:"title"`) to tell the JSON decoder how to
// map the JSON keys to our Go struct fields.
type Film struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Director    string `json:"director"`
	ReleaseDate string `json:"release_date"`
}

func main() {
	fmt.Println("Fetching films from Studio Ghibli API...")

	// 1. Make the HTTP GET request
	resp, err := http.Get("https://ghibliapi.vercel.app/films")
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	// Always defer closing the response body to prevent memory leaks!
	defer resp.Body.Close()

	// 2. Check the status code (we expect 200 OK)
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Expected status 200, got %d", resp.StatusCode)
	}

	// 3. Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// 4. Parse (Unmarshal) the JSON data into a slice of Film structs
	var films []Film
	if err := json.Unmarshal(body, &films); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	// 5. Print the results
	fmt.Printf("Successfully loaded %d films!\n\n", len(films))

	// Let's print out the first 3 films
	limit := min(len(films), 3)
	for i := range limit {
		film := films[i]
		fmt.Printf("Title: %s (%s)\n", film.Title, film.ReleaseDate)
		fmt.Printf("Director: %s\n", film.Director)
		fmt.Println("---")
	}
}
