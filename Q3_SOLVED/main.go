package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"unicode"
)

type MeatCounts map[string]int

func main() {
	http.HandleFunc("/beef/summary", beefSummaryHandler)

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func beefSummaryHandler(w http.ResponseWriter, r *http.Request) {
	textURL := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
	res, err := http.Get(textURL)
	if err != nil {
		http.Error(w, "Failed to fetch data", 500)
		return
	}
	defer res.Body.Close()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", 500)
		return
	}

	meats := handleText(string(body))

	// send response
	jsonData, err := json.Marshal(map[string]MeatCounts{"beef": meats})
	if err != nil {
		http.Error(w, "Failed to marshal JSON", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func handleText(text string) MeatCounts {
	counts := make(MeatCounts)

	// filter isSpace ด้วย because they have double space when filter ,.
	words := strings.FieldsFunc(text, func(r rune) bool {
		return strings.ContainsRune(",.", r) || unicode.IsSpace(r)
	})

	// fmt.Println(words)

	for _, word := range words {

		counts[strings.ToLower(word)]++
	}

	return counts
}
