package presentation

import (
	"fmt"
	"os/exec"
	"strings"
	"git-presenter/internal/slide"
)

type Presentation struct {
	branch       string
	slides       []*slide.Slide
	currentSlide *slide.Slide
}

func NewPresentation(config map[string]interface{}) *Presentation {
	branch, _ := config["branch"].(string)
	
	var slides []*slide.Slide
	if slidesData, ok := config["slides"].([]interface{}); ok {
		for _, slideData := range slidesData {
			if slideMap, ok := slideData.(map[string]interface{}); ok {
				if slideInfo, ok := slideMap["slide"].(map[string]interface{}); ok {
					slides = append(slides, slide.NewSlide(slideInfo))
				}
			}
		}
	}
	
	presentation := &Presentation{
		branch: branch,
		slides: slides,
	}
	
	// Find current slide based on current git HEAD
	presentation.currentSlide = presentation.findCurrentSlide()
	
	return presentation
}

func (p *Presentation) findCurrentSlide() *slide.Slide {
	cmd := exec.Command("git", "rev-parse", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		if len(p.slides) > 0 {
			return p.slides[0]
		}
		return nil
	}
	
	sha := strings.TrimSpace(string(output))
	for _, slide := range p.slides {
		if slide.Commit() == sha {
			return slide
		}
	}
	
	if len(p.slides) > 0 {
		return p.slides[0]
	}
	return nil
}

func (p *Presentation) Execute(command string) string {
	switch command {
	case "next", "n":
		return p.next()
	case "back", "b":
		return p.previous()
	case "start", "s":
		return p.start()
	case "end", "e":
		return p.end()
	case "list", "l":
		return p.list()
	case "help", "h":
		return p.help()
	case "exit":
		return p.exit()
	default:
		return "I canny understand ye, gonna try again"
	}
}

func (p *Presentation) start() string {
	if len(p.slides) == 0 {
		return "No slides available"
	}
	p.currentSlide = p.slides[0]
	return p.currentSlide.Execute()
}

func (p *Presentation) next() string {
	if len(p.slides) == 0 {
		return "No slides available"
	}
	
	currentIndex := p.getCurrentSlideIndex()
	if currentIndex == -1 || currentIndex >= len(p.slides)-1 {
		return p.currentSlide.Execute()
	}
	
	p.currentSlide = p.slides[currentIndex+1]
	return p.currentSlide.Execute()
}

func (p *Presentation) previous() string {
	if len(p.slides) == 0 {
		return "No slides available"
	}
	
	currentIndex := p.getCurrentSlideIndex()
	if currentIndex <= 0 {
		return p.currentSlide.Execute()
	}
	
	p.currentSlide = p.slides[currentIndex-1]
	return p.currentSlide.Execute()
}

func (p *Presentation) end() string {
	if len(p.slides) == 0 {
		return "No slides available"
	}
	p.currentSlide = p.slides[len(p.slides)-1]
	return p.currentSlide.Execute()
}

func (p *Presentation) list() string {
	if len(p.slides) == 0 {
		return "No slides available"
	}
	
	var result strings.Builder
	for _, slide := range p.slides {
		if slide == p.currentSlide {
			result.WriteString("*")
		}
		result.WriteString(slide.String())
		result.WriteString("\n")
	}
	return strings.TrimSpace(result.String())
}

func (p *Presentation) help() string {
	return `Git Presenter Reference

next/n: move to next slide
back/b: move back a slide
end/e:  move to end of presentation
start/s: move to start of presentation
list/l : list slides in presentation
help/h: display this message
!(exclimation mark): execute following in terminal
exit: exit from the presentation`
}

func (p *Presentation) exit() string {
	cmd := exec.Command("git", "checkout", "-q", p.branch)
	cmd.Run()
	return "Exited presentation"
}

func (p *Presentation) getCurrentSlideIndex() int {
	for i, slide := range p.slides {
		if slide == p.currentSlide {
			return i
		}
	}
	return -1
}

func (p *Presentation) StatusLine() string {
	if len(p.slides) == 0 {
		return "0/0 >"
	}
	position := p.getCurrentSlideIndex() + 1
	return fmt.Sprintf("%d/%d >", position, len(p.slides))
}