package internal

import (
	"testing"
)

func TestShouldCreatePresentationFromConfig(t *testing.T) {
	config := map[string]interface{}{
		"branch": "main",
		"slides": []interface{}{
			map[string]interface{}{
				"slide": map[string]interface{}{
					"commit":  "abc123",
					"message": "First commit",
				},
			},
		},
	}

	presentation := NewPresentation(config)

	if presentation == nil {
		t.Error("Expected presentation to be created")
	}
}
