package cricklib_test

import (
	"strings"
	"testing"
	"github.com/prajjwalyd/crickCLI/internal/cricklib"
)

func TestParseResponse(t *testing.T) {
	// Mock JSON response
	jsonResponse := `{"title": "Test Match", "update": "Inning 1", "livescore": "100/2", "runrate": "5.0"}`

	// Convert string to io.Reader
	reader := strings.NewReader(jsonResponse)

	// Test ParseResponse function
	matchData, err := cricklib.ParseResponse(reader)
	if err != nil {
		t.Errorf("Error parsing response: %v", err)
	}

	// Validate parsed data
	if matchData.Title != "Test Match" || matchData.Update != "Inning 1" {
		t.Errorf("Invalid parsed data: %v", matchData)
	}
	// Add more validation checks as needed
}
