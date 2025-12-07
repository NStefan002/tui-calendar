package main

import (
	"fmt"
	"log"
	"os"

	"github.com/NStefan002/tui-calendar/google"
	"github.com/NStefan002/tui-calendar/models"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// detect "init" command
	if len(os.Args) > 1 && os.Args[1] == "init" {
		runInit()
		return
	}

	// normal app launch
	f, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

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

func runInit() {
	cfgDir, err := google.AppConfigDir()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	credPath := cfgDir + "/credentials.json"

	if _, err := os.Stat(credPath); err == nil {
		fmt.Println("Setup already completed! You can now run:")
		fmt.Println("  tui-calendar")
		return
	}

	fmt.Println(`Welcome to tui-calendar setup instructions!

Before using Google Calendar, you must generate your own OAuth credentials.

NOTE: :: You will have to go through these steps only once. Your credentials and token
will be saved locally. Do not share them with anyone! ::

Steps:

1. Create a Google Cloud project:
   - Open: https://console.cloud.google.com/
   - Click the project dropdown (top-left)
   - Click "New Project", give it any name, and press "Create"

2. Enable the Google Calendar API for your project:
    - Open Navigation Menu (top-left) > APIs & Services > Enabled APIs & Services
    - Click "+ Enable APIs and Services" at the top
    - Search for "Google Calendar API"
    - Click it, then click "Enable"

3. Create OAuth 2.0 credentials:
    - Open Navigation Menu (top-left) > APIs & Services > Credentials
    - Click "Create Credentials" > "OAuth client ID"
    - If prompted to configure the consent screen, do so (choose "External", fill in required fields, and save)
    - Choose "Desktop app" as the Application type
    - Name it as you like and click "Create"
    - IMPORTANT: Download the credentials by clicking on 'Download JSON'
    - Click "OK" to exit the dialog

4. Save the downloaded file as:

   ` + credPath + `

5. Add yourself as a test user:
    - Open Navigation Menu (top-left) > APIs & Services > OAuth consent screen > Audience
    - Scroll down to "Test users" and click "Add users"
    - Add your Google account email and save

6. All done! You can now run:
    tui-calendar
[You will need to log in (one time action), and then you are finally free to use the app!]
`)
}
