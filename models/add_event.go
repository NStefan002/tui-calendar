// TODO: implement all fields

package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"

	"tui-calendar/styles"
	"tui-calendar/utils"
)

type addEventModel struct {
	title       textinput.Model // title of the new event
	description textarea.Model  // description of the new event
	startTime   time.Time       // start time of the new event
	endTime     time.Time       // end time of the new event
	location    textinput.Model // location of the new event
	allDay      bool            // whether the event is an all-day event
	idx         int             // index of the field being edited
}

func newAM() *addEventModel {
	am := &addEventModel{
		title:       textinput.New(),
		description: textarea.New(),
		location:    textinput.New(),
		startTime:   time.Now(),
		endTime:     time.Now().Add(time.Hour), // default to 1 hour later
		allDay:      false,
		idx:         0,
	}
	am.title.Placeholder = "Title"
	am.title.Width = 50
	am.title.CharLimit = 50
	am.title.Focus()
	am.title.PromptStyle = styles.ActiveTextinput
	am.title.TextStyle = styles.ActiveTextinput
	am.title.Cursor.Style = styles.ActiveTextinput

	am.description.Placeholder = "Description"
	am.description.FocusedStyle.Prompt = styles.ActiveTextinput
	am.description.FocusedStyle.Text = styles.ActiveTextinput
	am.description.BlurredStyle.Prompt = styles.InactiveTextinput
	am.description.BlurredStyle.Text = styles.InactiveTextinput

	am.location.Placeholder = "Location"
	am.location.Width = 50
	am.location.CharLimit = 100
	am.location.PromptStyle = styles.InactiveTextinput
	am.location.TextStyle = styles.InactiveTextinput
	am.location.Cursor.Style = styles.InactiveTextinput

	return am
}

func (am *addEventModel) view(m *model) string {
	var sb strings.Builder

	// header
	header := styles.Header.Render(fmt.Sprintf("➕ Add Event for %s", m.cm.selected.Format("January 2, 2006")))
	sb.WriteString(utils.CenterText(header, m.screenWidth) + "\n\n")

	// form fields
	fields := []string{
		lipgloss.JoinHorizontal(lipgloss.Top, styles.FieldLabel.Render("Title:"), am.title.View()),
		"\n",
		lipgloss.JoinHorizontal(lipgloss.Top, styles.FieldLabel.Render("Description:"), am.description.View()),
		"\n",
		lipgloss.JoinHorizontal(lipgloss.Top, styles.FieldLabel.Render("Location:"), am.location.View()),
	}

	form := lipgloss.JoinVertical(lipgloss.Left, fields...)
	box := styles.Box.Render(form)

	sb.WriteString(utils.CenterText(box, m.screenWidth))

	helpText := m.help.View(m.addEventViewKeys)
	sb.WriteString("\n\n" + utils.CenterText(helpText, m.screenWidth))

	return sb.String()
}

func (am *addEventModel) changeFocus(direction int) {
	switch am.idx {
	case 0:
		am.title.Blur()
		am.title.PromptStyle = styles.InactiveTextinput
		am.title.TextStyle = styles.InactiveTextinput
		am.title.Cursor.Style = styles.InactiveTextinput
	case 1:
		am.description.Blur()
	case 2:
	case 3:
	case 4:
		am.location.Blur()
		am.location.PromptStyle = styles.InactiveTextinput
		am.location.TextStyle = styles.InactiveTextinput
		am.location.Cursor.Style = styles.InactiveTextinput
	case 5:
	}

	am.idx = (am.idx + direction) % 6

	switch am.idx {
	case 0:
		am.title.Focus()
		am.title.PromptStyle = styles.ActiveTextinput
		am.title.TextStyle = styles.ActiveTextinput
		am.title.Cursor.Style = styles.ActiveTextinput
	case 1:
		am.description.Focus()
	case 2:
	case 3:
	case 4:
		am.location.Focus()
		am.location.PromptStyle = styles.ActiveTextinput
		am.location.TextStyle = styles.ActiveTextinput
		am.location.Cursor.Style = styles.ActiveTextinput
	case 5:
	}
}
