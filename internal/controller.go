package internal

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"gopkg.in/yaml.v2"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Controller struct {
	currentDir string
}

type PresentationConfig struct {
	Slides []SlideConfig `yaml:"slides"`
	Branch string        `yaml:"branch"`
}

type SlideConfig struct {
	Slide map[string]interface{} `yaml:"slide"`
}

func NewController(currentDir string) *Controller {
	return &Controller{
		currentDir: currentDir,
	}
}

func (c *Controller) InitializePresentation() error {
	slides, err := c.createSlides()
	if err != nil {
		return err
	}

	branch, err := c.getCurrentBranch()
	if err != nil {
		return err
	}

	config := PresentationConfig{
		Slides: slides,
		Branch: branch,
	}

	yamlData, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	configFile := filepath.Join(c.currentDir, ".presentation")
	err = os.WriteFile(configFile, yamlData, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Presentation has been initialised")
	fmt.Println("run 'git-presenter start' to begin the presentation")

	return nil
}

func (c *Controller) createSlides() ([]SlideConfig, error) {
	repo, err := git.PlainOpen(c.currentDir)
	if err != nil {
		return nil, err
	}

	ref, err := repo.Head()
	if err != nil {
		return nil, err
	}

	cIter, err := repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		return nil, err
	}

	var slides []SlideConfig
	err = cIter.ForEach(func(c *object.Commit) error {
		slideData := map[string]interface{}{
			"commit":  c.Hash.String(),
			"message": c.Message,
		}
		slides = append(slides, SlideConfig{Slide: slideData})
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Reverse slides to match Ruby implementation (oldest first)
	for i, j := 0, len(slides)-1; i < j; i, j = i+1, j-1 {
		slides[i], slides[j] = slides[j], slides[i]
	}

	return slides, nil
}

func (c *Controller) getCurrentBranch() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Dir = c.currentDir
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func (c *Controller) StartPresentation() (*PresentationConfig, error) {
	configFile := filepath.Join(c.currentDir, ".presentation")
	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	var config PresentationConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
