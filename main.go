package main

import (
	"log"
	"os"

	"github.com/NStefan002/tui-calendar/google"
	"github.com/NStefan002/tui-calendar/models"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	// open or create log file
	f, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// redirect logger to the file
	log.SetOutput(f)

	srv, err := google.GetClient()
	if err != nil {
		log.Fatalf("Failed to get calendar client: %v", err)
	}

	p := tea.NewProgram(models.CreateModel(srv))
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
