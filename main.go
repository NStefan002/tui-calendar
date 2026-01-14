package main

import (
	"log"
	"os"

	"github.com/NStefan002/tui-calendar/v2/google"
	"github.com/NStefan002/tui-calendar/v2/models"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// open log file
	log_file := os.TempDir() + "/tui-calendar.log"
	f, err := os.OpenFile(log_file, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	log.SetOutput(f)

	// detect "init" command
	if len(os.Args) > 1 && os.Args[1] == "init" {
		runInit()
		return
	}

	srv, err := google.GetClient()
	if err != nil {
		log.Fatalf("Failed to get calendar client: %v", err)
	}

	p := tea.NewProgram(models.CreateModel(srv))
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func runInit() {
	p := tea.NewProgram(models.CreateInitModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
