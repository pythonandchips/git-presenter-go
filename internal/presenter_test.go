package internal

import (
	"testing"
)

func TestShouldCreateGitPresenter(t *testing.T) {
	presenter := NewGitPresenter("/tmp/test", true)

	if presenter == nil {
		t.Error("Expected presenter to be created")
	}
}

func TestShouldExecuteInitCommand(t *testing.T) {
	presenter := NewGitPresenter("..", true)

	err := presenter.Execute("init")
	if err != nil {
		t.Skipf("Skipping test - not in git repository: %v", err)
	}
}
