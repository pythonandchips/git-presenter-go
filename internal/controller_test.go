package internal

import (
	"testing"
)

func TestShouldCreateController(t *testing.T) {
	controller := NewController("/tmp/test")

	if controller == nil {
		t.Error("Expected controller to be created")
	}
}

func TestShouldInitializePresentation(t *testing.T) {
	controller := NewController("..")

	err := controller.InitializePresentation()
	if err != nil {
		t.Skipf("Skipping test - not in git repository: %v", err)
	}
}
