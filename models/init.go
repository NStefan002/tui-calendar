package models

import (
	"log"
	"os"

	"github.com/NStefan002/tui-calendar/v2/google"
	"github.com/NStefan002/tui-calendar/v2/styles"
	"github.com/NStefan002/tui-calendar/v2/utils"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"

	tea "github.com/charmbracelet/bubbletea"
)

type initModel struct {
	viewport viewport.Model
	help     help.Model

	ready bool

	keys initViewKeyMap // key bindings

	screenWidth  int // width of the terminal screen
	screenHeight int // height of the terminal screen

	errMessage string // error message to display, if any
}

func CreateInitModel() initModel {
	m := initModel{
		help: help.New(),

		ready: false,

		keys: initViewKeys,
	}
	return m
}

func (m initModel) Init() tea.Cmd {
	return nil
}

func (m initModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.screenWidth = msg.Width
		m.screenHeight = msg.Height
		headerHeight := lipgloss.Height(m.getHeader())
		footerHeight := lipgloss.Height(m.getFooter())
		verticalMarginHeight := headerHeight + footerHeight

		content := m.getContent()
		if !m.ready {
			contentHeight := lipgloss.Height(content)
			m.viewport = viewport.New(m.screenWidth, min(contentHeight, m.screenHeight-verticalMarginHeight))
			// m.viewport.YPosition = headerHeight
			m.viewport.SetContent(m.getContent())
			m.ready = true
		} else {
			m.viewport.Width = m.screenWidth
			m.viewport.Height = min(lipgloss.Height(m.getContent()), msg.Height-verticalMarginHeight)
		}

	case errMsg:
		m.errMessage = msg.Error()

	case tea.KeyMsg:
		// clear error message on any key press
		if m.errMessage != "" {
			m.errMessage = ""
			return m, nil
		}
		if key.Matches(msg, m.keys.Quit) {
			return m, tea.Quit
		}
	}
	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m initModel) View() string {
	if m.errMessage != "" {
		return errorView(m.errMessage, m.screenWidth)
	}
	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.getHeader(),
		m.viewport.View(),
		m.getFooter(),
	)
}

func (m initModel) getHeader() string {
	return utils.CenterText(styles.InitTitle.Render("Welcome to tui-calendar setup instructions!"), m.screenWidth)
}

func (m initModel) getFooter() string {
	return utils.CenterText(m.help.View(m.keys), m.screenWidth)
}

func (m initModel) getContent() string {
	credPath, err := google.CredentialsFilePath()
	if err != nil {
		log.Println("Error getting credentials file path:", err)
	}

	if _, err := os.Stat(credPath); err == nil {
		content := lipgloss.JoinVertical(
			lipgloss.Left,
			styles.InitStep.Render("Setup already completed, your credentials are stored at:"),
			styles.InitPath.Render(credPath),
			"",
			styles.InitStep.Render("You can now run: ")+styles.InitHint.Render("tui-calendar")+styles.InitStep.Render(" to start the application!"),
		)
		box := styles.InitBox.Render(content)
		return utils.CenterText(box, m.screenWidth)
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		styles.InitStep.Render("Before using Google Calendar, you must generate your own OAuth credentials."),
		"",
		styles.InitStep.Render("You will have to go through these steps only once. Your credentials and token will be saved locally. ")+styles.InitPath.Render("Do not share them with anyone!"),
		"",
		"",
		styles.InitStep.Render("1. Create a Google Cloud project"),
		styles.InitSubStep.Render("  - Open: ")+styles.InitPath.Render("https://console.cloud.google.com/"),
		styles.InitSubStep.Render("  - Click the project dropdown (top-left)"),
		styles.InitSubStep.Render("  - Click ")+styles.InitPath.Render("New Project")+styles.InitSubStep.Render(", give it any name, and press ")+styles.InitPath.Render("Create"),
		"",
		styles.InitStep.Render("2. Enable Google Calendar API"),
		styles.InitSubStep.Render("  - Open Navigation Menu (top-left) > APIs & Services > Enabled APIs & Services"),
		styles.InitSubStep.Render("  - Click ")+styles.InitPath.Render("+ Enable APIs and Services")+styles.InitSubStep.Render(" at the top"),
		styles.InitSubStep.Render("  - Search for ")+styles.InitPath.Render("Google Calendar API"),
		styles.InitSubStep.Render("  - Click it, then click ")+styles.InitPath.Render("Enable"),
		"",
		styles.InitStep.Render("3. Create OAuth Client ID (Desktop app)"),

		styles.InitSubStep.Render("  - Choose ")+styles.InitPath.Render("Desktop app")+styles.InitSubStep.Render(" as the Application type"),
		styles.InitSubStep.Render("  - Name it as you like and click ")+styles.InitPath.Render("Create"),
		styles.InitSubStep.Render("  - IMPORTANT: Download the credentials by clicking on 'Download JSON'"),
		styles.InitSubStep.Render("  - Click ")+styles.InitPath.Render("OK")+styles.InitSubStep.Render(" to exit the dialog"),
		"",
		styles.InitStep.Render("4. Save the downloaded file as:"),
		styles.InitPath.Render(credPath),
		"",
		styles.InitStep.Render("5. Add yourself as a test user"),
		styles.InitSubStep.Render("  - Open Navigation Menu (top-left) > APIs & Services > OAuth consent screen > Audience"),
		styles.InitSubStep.Render("  - Scroll down to ")+styles.InitPath.Render("Test users")+styles.InitSubStep.Render(" and click ")+styles.InitPath.Render("Add users"),
		styles.InitSubStep.Render("  - Add your Google account email and save"),
		"",
		styles.InitStep.Render("6. All done! Run: ")+styles.InitHint.Render("tui-calendar")+styles.InitStep.Render(", log in (one time action), and enjoy your terminal calendar!"),
	)

	box := styles.InitBox.Render(content)

	return utils.CenterText(box, m.screenWidth)
}
