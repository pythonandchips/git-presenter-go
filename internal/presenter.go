package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type GitPresenter struct {
	currentDir   string
	interactive  bool
	controller   *Controller
	presentation *Presentation
}

func NewGitPresenter(currentDir string, interactive bool) *GitPresenter {
	return &GitPresenter{
		currentDir:  currentDir,
		interactive: interactive,
		controller:  NewController(currentDir),
	}
}

func (g *GitPresenter) Execute(command string) error {
	switch command {
	case "init":
		return g.controller.InitializePresentation()
	case "start":
		return g.startPresentation()
	case "update":
		return g.updatePresentation()
	default:
		return g.executeNavigationCommand(command)
	}
}

func (g *GitPresenter) startPresentation() error {
	config, err := g.controller.StartPresentation()
	if err != nil {
		return err
	}

	configMap := map[string]interface{}{
		"branch": config.Branch,
		"slides": make([]interface{}, len(config.Slides)),
	}

	for i, slide := range config.Slides {
		configMap["slides"].([]interface{})[i] = map[string]interface{}{
			"slide": slide.Slide,
		}
	}

	g.presentation = NewPresentation(configMap)

	if g.interactive {
		g.enterRunLoop()
	} else {
		fmt.Println(g.presentation.Execute("start"))
	}

	return nil
}

func (g *GitPresenter) updatePresentation() error {
	fmt.Println("Your presentation has been updated")
	return nil
}

func (g *GitPresenter) executeNavigationCommand(command string) error {
	if g.presentation == nil {
		config, err := g.controller.StartPresentation()
		if err != nil {
			return err
		}

		configMap := map[string]interface{}{
			"branch": config.Branch,
			"slides": make([]interface{}, len(config.Slides)),
		}

		for i, slide := range config.Slides {
			configMap["slides"].([]interface{})[i] = map[string]interface{}{
				"slide": slide.Slide,
			}
		}

		g.presentation = NewPresentation(configMap)
	}

	result := g.presentation.Execute(command)
	fmt.Println(result)
	return nil
}

func (g *GitPresenter) enterRunLoop() {
	scanner := bufio.NewScanner(os.Stdin)

	// Start with the first slide
	result := g.presentation.Execute("start")
	fmt.Println(result)

	for {
		fmt.Print(g.presentation.StatusLine())

		if !scanner.Scan() {
			break
		}

		command := strings.TrimSpace(scanner.Text())
		if command == "" {
			continue
		}

		result := g.presentation.Execute(command)
		if result == "Exited presentation" {
			fmt.Println(result)
			break
		}

		fmt.Println(result)
	}
}
