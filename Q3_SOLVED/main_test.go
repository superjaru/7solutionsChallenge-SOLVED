package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBeefSummaryHandler(t *testing.T) {
	// Create a mock HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(beefSummaryHandler))
	defer mockServer.Close()

	// Send a GET request to the mock server
	resp, err := http.Get(mockServer.URL + "/beef/summary")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	// checks for the response body
	var responseBody map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		t.Fatalf("Error decoding JSON: %v", err)
	}

	// Check beef field
	if _, ok := responseBody["beef"]; !ok {
		t.Error("Expected 'beef' field in the response")
	}
}

func TestHandleText(t *testing.T) {
	// Test cases
	tests := []struct {
		input  string
		output map[string]int
	}{
		{"test,, lorem ipsum   .test. labtest", map[string]int{"test": 2, "lorem": 1, "ipsum": 1, "labtest": 1}},
		{"", map[string]int{}}, // Test case for empty input

	}

	for _, test := range tests {
		result := handleText(test.input)
		// Check if the result matches the expected output
		if !equal(result, test.output) {
			t.Errorf("handleText(%q) returned %v, expected %v", test.input, result, test.output)
		}
	}
}

// Helper function to check if two maps are equal
func equal(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}
