package slide

import (
	"os/exec"
	"strings"
)

type Slide struct {
	commit  string
	message string
}

func NewSlide(slideData map[string]interface{}) *Slide {
	commit, _ := slideData["commit"].(string)
	message, _ := slideData["message"].(string)
	
	return &Slide{
		commit:  commit,
		message: message,
	}
}

func (s *Slide) Commit() string {
	return s.commit
}

func (s *Slide) Message() string {
	return s.message
}

func (s *Slide) Execute() string {
	if s.commit == "" {
		return ""
	}
	
	// Execute git checkout commands
	cmd := exec.Command("git", "checkout", "-q", ".")
	cmd.Run()
	
	cmd = exec.Command("git", "checkout", "-q", s.commit)
	cmd.Run()
	
	return strings.TrimSpace(s.message)
}

func (s *Slide) String() string {
	if s.commit == "" {
		return ""
	}
	commitShort := s.commit
	if len(commitShort) > 10 {
		commitShort = commitShort[:10]
	}
	return commitShort + ", " + strings.TrimSpace(s.message)
}