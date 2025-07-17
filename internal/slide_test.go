package internal

import (
	"testing"
)

func TestShouldCreateSlideFromCommitData(t *testing.T) {
	slideData := map[string]interface{}{
		"commit":  "abc123",
		"message": "Initial commit",
	}

	slide := NewSlide(slideData)

	if slide.Commit() != "abc123" {
		t.Errorf("Expected commit 'abc123', got '%s'", slide.Commit())
	}

	if slide.Message() != "Initial commit" {
		t.Errorf("Expected message 'Initial commit', got '%s'", slide.Message())
	}
}
