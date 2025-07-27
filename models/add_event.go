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

	am.location.Placeholder = "Location"
	am.location.Width = 50
	am.location.CharLimit = 100
	am.location.PromptStyle = styles.InactiveTextinput
	am.location.TextStyle = styles.InactiveTextinput
	am.location.Cursor.Style = styles.InactiveTextinput

	return am
}

func (am *addEventModel) view(selectedDate time.Time, scrWidth, scrHeight int) string {
	var sb strings.Builder

	// header
	header := styles.Header.Render(fmt.Sprintf("➕ Add Event for %s", selectedDate.Format("January 2, 2006")))
	sb.WriteString(utils.CenterText(header, scrWidth) + "\n\n")

	// form fields
	fields := []string{
		lipgloss.JoinHorizontal(lipgloss.Top, styles.FieldLabel.Render("Title:"), am.title.View()),
		lipgloss.JoinHorizontal(lipgloss.Top, styles.FieldLabel.Render("Location:"), am.location.View()),
		// you can add more like description, start time, end time similarly.
	}

	form := lipgloss.JoinVertical(lipgloss.Left, fields...)
	box := styles.Box.Render(form)

	sb.WriteString(utils.CenterText(box, scrWidth))

	// footer
	footer := styles.FormFooter.Render("[tab] Next field  [enter] Confirm  [esc/q] Cancel")
	sb.WriteString("\n\n" + utils.CenterText(footer, scrWidth))

	return sb.String()
}

func (am *addEventModel) nextField() {
	if am.title.Focused() {
		am.title.Blur()
		am.title.PromptStyle = styles.InactiveTextinput
		am.title.TextStyle = styles.InactiveTextinput
		am.title.Cursor.Style = styles.InactiveTextinput

		am.location.Focus()
		am.location.PromptStyle = styles.ActiveTextinput
		am.location.TextStyle = styles.ActiveTextinput
		am.location.Cursor.Style = styles.ActiveTextinput
	} else if am.location.Focused() {
		am.location.Blur()
		am.location.PromptStyle = styles.InactiveTextinput
		am.location.TextStyle = styles.InactiveTextinput
		am.location.Cursor.Style = styles.InactiveTextinput

		am.title.Focus()
		am.title.PromptStyle = styles.ActiveTextinput
		am.title.TextStyle = styles.ActiveTextinput
		am.title.Cursor.Style = styles.ActiveTextinput
	}
}

func (am *addEventModel) prevField() {
	if am.title.Focused() {
		am.title.Blur()
		am.title.PromptStyle = styles.InactiveTextinput
		am.title.TextStyle = styles.InactiveTextinput
		am.title.Cursor.Style = styles.InactiveTextinput

		am.location.Focus()
		am.location.PromptStyle = styles.ActiveTextinput
		am.location.TextStyle = styles.ActiveTextinput
		am.location.Cursor.Style = styles.ActiveTextinput
	} else if am.location.Focused() {
		am.location.Blur()
		am.location.PromptStyle = styles.InactiveTextinput
		am.location.TextStyle = styles.InactiveTextinput
		am.location.Cursor.Style = styles.InactiveTextinput

		am.title.Focus()
		am.title.PromptStyle = styles.ActiveTextinput
		am.title.TextStyle = styles.ActiveTextinput
		am.title.Cursor.Style = styles.ActiveTextinput
	}
}
