package models

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
	"google.golang.org/api/calendar/v3"

	"github.com/NStefan002/tui-calendar/v2/styles"
	"github.com/NStefan002/tui-calendar/v2/utils"
)

type repeatOption int

const (
	No repeatOption = iota
	Daily
	Weekly
	Monthly
	Yearly
)

type addEventModel struct {
	titleInput       textinput.Model // 0
	descriptionInput textarea.Model  // 1
	startTimeInput   textinput.Model // 2
	endTimeInput     textinput.Model // 3
	locationInput    textinput.Model // 4
	allDayInput      textinput.Model // 5
	repeatInput      textinput.Model // 6

	title       string
	description string
	startTime   time.Time
	endTime     time.Time
	location    string
	allDay      bool
	repeat      repeatOption

	selectedDate time.Time

	idx int
}

func newAM() *addEventModel {
	am := &addEventModel{
		titleInput:       textinput.New(),
		descriptionInput: textarea.New(),
		startTimeInput:   textinput.New(),
		endTimeInput:     textinput.New(),
		locationInput:    textinput.New(),
		allDayInput:      textinput.New(),
		repeatInput:      textinput.New(),

		title:       "",
		description: "",
		location:    "",
		allDay:      false,
		repeat:      No,

		idx: 0,
	}

	// --- Title ---
	am.titleInput.Placeholder = "Title"
	am.titleInput.CharLimit = 50
	am.titleInput.Width = 50
	am.titleInput.Focus()
	am.titleInput.PromptStyle = styles.ActiveTextinput
	am.titleInput.TextStyle = styles.ActiveTextinput
	am.titleInput.Cursor.Style = styles.ActiveTextinput

	// --- Description ---
	am.descriptionInput.Placeholder = "Description"
	am.descriptionInput.FocusedStyle.Prompt = styles.ActiveTextinput
	am.descriptionInput.FocusedStyle.Text = styles.ActiveTextinput
	am.descriptionInput.BlurredStyle.Prompt = styles.InactiveTextinput
	am.descriptionInput.BlurredStyle.Text = styles.InactiveTextinput

	// --- Start time ---
	am.startTimeInput.Placeholder = "HH:MM"
	am.startTimeInput.Width = 8
	am.startTimeInput.PromptStyle = styles.InactiveTextinput
	am.startTimeInput.TextStyle = styles.InactiveTextinput
	am.startTimeInput.Cursor.Style = styles.InactiveTextinput

	// --- End time ---
	am.endTimeInput.Placeholder = "HH:MM"
	am.endTimeInput.Width = 8
	am.endTimeInput.PromptStyle = styles.InactiveTextinput
	am.endTimeInput.TextStyle = styles.InactiveTextinput
	am.endTimeInput.Cursor.Style = styles.InactiveTextinput

	am.initTimes()

	// --- Location ---
	am.locationInput.Placeholder = "Location"
	am.locationInput.Width = 50
	am.locationInput.CharLimit = 100
	am.locationInput.PromptStyle = styles.InactiveTextinput
	am.locationInput.TextStyle = styles.InactiveTextinput
	am.locationInput.Cursor.Style = styles.InactiveTextinput

	// --- All-day ---
	am.allDayInput.Prompt = ""
	am.allDayInput.SetValue("   ")
	am.allDayInput.PromptStyle = styles.InactiveTextinput
	am.allDayInput.TextStyle = styles.InactiveTextinput

	// --- Repeat ---
	am.repeatInput.Prompt = ""
	am.repeatInput.SetValue("[no] |  daily  |  weekly  |  monthly  |  yearly  ")
	am.repeatInput.PromptStyle = styles.InactiveTextinput
	am.repeatInput.TextStyle = styles.InactiveTextinput

	return am
}

func (am *addEventModel) view(m *model) string {
	var sb strings.Builder

	// Header
	header := styles.Header.Render(fmt.Sprintf("➕ Add Event for %s", m.cm.selected.Format("January 2, 2006")))
	sb.WriteString(utils.CenterText(header, m.screenWidth) + "\n\n")

	formFields := []string{
		lipgloss.JoinHorizontal(lipgloss.Top, styles.FieldLabel.Render("Title:"), am.titleInput.View()),
		"\n",
		lipgloss.JoinHorizontal(lipgloss.Top, styles.FieldLabel.Render("Description:"), am.descriptionInput.View()),
		"\n",
		lipgloss.JoinHorizontal(lipgloss.Top, styles.FieldLabel.Render("Start:"), am.startTimeInput.View()),
		"\n",
		lipgloss.JoinHorizontal(lipgloss.Top, styles.FieldLabel.Render("End:"), am.endTimeInput.View()),
		"\n",
		lipgloss.JoinHorizontal(lipgloss.Top, styles.FieldLabel.Render("Location:"), am.locationInput.View()),
		"\n",
		lipgloss.JoinHorizontal(lipgloss.Top, styles.FieldLabel.Render("All-day:"), am.allDayInput.View()),
		"\n",
		lipgloss.JoinHorizontal(lipgloss.Top, styles.FieldLabel.Render("Repeat:"), am.repeatInput.View()),
	}

	form := lipgloss.JoinVertical(lipgloss.Left, formFields...)
	box := styles.Box.
		Border(lipgloss.RoundedBorder()).
		Padding(1, 2).
		Render(form)

	sb.WriteString(utils.CenterText(box, m.screenWidth))

	// help text
	help := m.help.View(m.addEventViewKeys)
	sb.WriteString("\n\n" + utils.CenterText(help, m.screenWidth))

	return sb.String()
}

func (am *addEventModel) checkBoxFocused() bool {
	return am.idx == 5
}

func (am *addEventModel) timeFieldFocused() bool {
	return am.idx == 2 || am.idx == 3
}

func (am *addEventModel) repeatFieldFocused() bool {
	return am.idx == 6
}

// Rotates focus between fields 0..5
func (am *addEventModel) changeFocus(direction int) {
	// blur current
	switch am.idx {
	case 0:
		am.titleInput.Blur()
		am.titleInput.PromptStyle = styles.InactiveTextinput
		am.titleInput.TextStyle = styles.InactiveTextinput
		am.titleInput.Cursor.Style = styles.InactiveTextinput
	case 1:
		am.descriptionInput.Blur()
	case 2:
		am.startTimeInput.Blur()
		am.startTimeInput.PromptStyle = styles.InactiveTextinput
		am.startTimeInput.TextStyle = styles.InactiveTextinput
		am.startTimeInput.Cursor.Style = styles.InactiveTextinput
	case 3:
		am.endTimeInput.Blur()
		am.endTimeInput.PromptStyle = styles.InactiveTextinput
		am.endTimeInput.TextStyle = styles.InactiveTextinput
		am.endTimeInput.Cursor.Style = styles.InactiveTextinput
	case 4:
		am.locationInput.Blur()
		am.locationInput.PromptStyle = styles.InactiveTextinput
		am.locationInput.TextStyle = styles.InactiveTextinput
		am.locationInput.Cursor.Style = styles.InactiveTextinput
	case 5:
		am.allDayInput.Blur()
		am.allDayInput.PromptStyle = styles.InactiveTextinput
		am.allDayInput.TextStyle = styles.InactiveTextinput
	case 6:
		am.repeatInput.Blur()
		am.repeatInput.PromptStyle = styles.InactiveTextinput
		am.repeatInput.TextStyle = styles.InactiveTextinput
	}

	am.idx = (am.idx + direction + 7) % 7

	// focus new
	switch am.idx {
	case 0:
		am.titleInput.Focus()
		am.titleInput.PromptStyle = styles.ActiveTextinput
		am.titleInput.TextStyle = styles.ActiveTextinput
		am.titleInput.Cursor.Style = styles.ActiveTextinput
	case 1:
		am.descriptionInput.Focus()
	case 2:
		am.startTimeInput.Focus()
		am.startTimeInput.PromptStyle = styles.ActiveTextinput
		am.startTimeInput.TextStyle = styles.ActiveTextinput
		am.startTimeInput.Cursor.Style = styles.ActiveTextinput
	case 3:
		am.endTimeInput.Focus()
		am.endTimeInput.PromptStyle = styles.ActiveTextinput
		am.endTimeInput.TextStyle = styles.ActiveTextinput
		am.endTimeInput.Cursor.Style = styles.ActiveTextinput
	case 4:
		am.locationInput.Focus()
		am.locationInput.PromptStyle = styles.ActiveTextinput
		am.locationInput.TextStyle = styles.ActiveTextinput
		am.locationInput.Cursor.Style = styles.ActiveTextinput
	case 5:
		am.allDayInput.Focus()
		am.allDayInput.PromptStyle = styles.ActiveTextinput
		am.allDayInput.TextStyle = styles.ActiveTextinput
	case 6:
		am.repeatInput.Focus()
		am.repeatInput.PromptStyle = styles.ActiveTextinput
		am.repeatInput.TextStyle = styles.ActiveTextinput
	}
}

func (am *addEventModel) resetForm() {
	am.titleInput.SetValue("")
	am.descriptionInput.SetValue("")
	am.locationInput.SetValue("")
	am.allDay = false
	am.allDayInput.SetValue("   ")
	am.repeat = No
	am.repeatInput.SetValue("[no] |  daily  |  weekly  |  monthly  |  yearly  ")
	am.initTimes()
	am.idx = 0
	am.title = ""
	am.description = ""
	am.location = ""
	am.changeFocus(0)
}

func (am *addEventModel) prefillForm(event *calendar.Event) {
	am.titleInput.SetValue(event.Summary)
	am.descriptionInput.SetValue(event.Description)
	am.locationInput.SetValue(event.Location)

	// check if all-day
	if event.Start.Date != "" && event.End.Date != "" {
		am.allDay = true
		am.allDayInput.SetValue("✓  ")
		am.startTimeInput.SetValue("     ")
		am.endTimeInput.SetValue("     ")
	} else {
		startTime, err := time.Parse(time.RFC3339, event.Start.DateTime)
		if err == nil {
			am.startTime = startTime
			am.startTimeInput.SetValue(am.startTime.Format("15:04"))
		}

		endTime, err := time.Parse(time.RFC3339, event.End.DateTime)
		if err == nil {
			am.endTime = endTime
			am.endTimeInput.SetValue(am.endTime.Format("15:04"))
		}
	}
}

func (am *addEventModel) initTimes() {
	now := time.Now()
	if am.allDay {
		am.startTime = am.selectedDate
		am.endTime = am.selectedDate.Add(24 * time.Hour)
		am.startTimeInput.SetValue(am.startTime.Format("     "))
		am.endTimeInput.SetValue(am.endTime.Format("     "))
	} else {
		am.startTime = time.Date(am.selectedDate.Year(), am.selectedDate.Month(), am.selectedDate.Day(), now.Hour(), now.Minute(), 0, 0, am.selectedDate.Location())
		am.endTime = am.startTime.Add(time.Hour)
		am.startTimeInput.SetValue(am.startTime.Format("15:04"))
		am.endTimeInput.SetValue(am.endTime.Format("15:04"))
	}
}

func (am *addEventModel) toggleAllDay() {
	if am.idx != 5 {
		return
	}

	am.allDay = !am.allDay
	am.initTimes()
	am.allDayInput.SetValue(func() string {
		if am.allDay {
			return "✓  "
		}
		return "   "
	}())
}

func (am *addEventModel) changeMinutes(delta int) {
	if am.idx == 2 || am.idx == 3 {
		var input *textinput.Model
		var currentTime time.Time

		if am.idx == 2 {
			input = &am.startTimeInput
			currentTime = am.startTime
		} else {
			input = &am.endTimeInput
			currentTime = am.endTime
		}

		// parse current input value
		parsedTime, err := time.Parse("15:04", input.Value())
		if err != nil {
			return // invalid time format, do nothing
		}

		// create a new time with the same date but updated time
		newTime := time.Date(
			currentTime.Year(),
			currentTime.Month(),
			currentTime.Day(),
			parsedTime.Hour(),
			parsedTime.Minute(),
			0, 0, currentTime.Location(),
		).Add(time.Duration(delta) * time.Minute)

		// update the input value
		input.SetValue(newTime.Format("15:04"))

		// update the corresponding field
		if am.idx == 2 {
			am.startTime = newTime
		} else {
			am.endTime = newTime
		}
	}
}

func (am *addEventModel) changeHours(delta int) {
	if am.idx == 2 || am.idx == 3 {
		var input *textinput.Model
		var currentTime time.Time

		if am.idx == 2 {
			input = &am.startTimeInput
			currentTime = am.startTime
		} else {
			input = &am.endTimeInput
			currentTime = am.endTime
		}

		// parse current input value
		parsedTime, err := time.Parse("15:04", input.Value())
		if err != nil {
			return // invalid time format, do nothing
		}

		// create a new time with the same date but updated time
		newTime := time.Date(
			currentTime.Year(),
			currentTime.Month(),
			currentTime.Day(),
			parsedTime.Hour(),
			parsedTime.Minute(),
			0, 0, currentTime.Location(),
		).Add(time.Duration(delta) * time.Hour)

		// update the input value
		input.SetValue(newTime.Format("15:04"))

		// update the corresponding field
		if am.idx == 2 {
			am.startTime = newTime
		} else {
			am.endTime = newTime
		}
	}
}

func (am *addEventModel) repeatToRRule() string {
	switch am.repeat {
	case Daily:
		return "RRULE:FREQ=DAILY"
	case Weekly:
		return "RRULE:FREQ=WEEKLY"
	case Monthly:
		return "RRULE:FREQ=MONTHLY"
	case Yearly:
		return "RRULE:FREQ=YEARLY"
	}
	return ""
}

func (am *addEventModel) updateRepeatOpts(inc int) {
	if am.idx != 6 {
		return
	}
	switch am.repeat {
	case No:
		if inc == 1 {
			am.repeat = Daily
		} else {
			am.repeat = Yearly
		}
	case Daily:
		if inc == 1 {
			am.repeat = Weekly
		} else {
			am.repeat = No
		}
	case Weekly:
		if inc == 1 {
			am.repeat = Monthly
		} else {
			am.repeat = Daily
		}
	case Monthly:
		if inc == 1 {
			am.repeat = Yearly
		} else {
			am.repeat = Weekly
		}
	case Yearly:
		if inc == 1 {
			am.repeat = No
		} else {
			am.repeat = Monthly
		}
	}

	// TODO: style nicer
	am.repeatInput.SetValue(func() string {
		log.Println(am.repeat)
		switch am.repeat {
		case Daily:
			return " no  | [daily] |  weekly  |  monthly  |  yearly  "
		case Weekly:
			return " no  |  daily  | [weekly] |  monthly  |  yearly  "
		case Monthly:
			return " no  |  daily  |  weekly  | [monthly] |  yearly  "
		case Yearly:
			return " no  |  daily  |  weekly  |  monthly  | [yearly] "
		default:
			return "[no] |  daily  |  weekly  |  monthly  |  yearly  "
		}
	}())
}

func (am *addEventModel) submitAddEventForm() (*calendar.Event, error) {
	am.title = am.titleInput.Value()
	am.description = am.descriptionInput.Value()
	am.location = am.locationInput.Value()

	rrule := am.repeatToRRule()
	hasRepeat := rrule != ""

	// handle all-day
	if am.allDay && !hasRepeat {
		// Valid all-day, because RRULE is NOT present
		return &calendar.Event{
			Summary:     am.title,
			Description: am.description,
			Location:    am.location,
			Start: &calendar.EventDateTime{
				Date: am.startTime.Format("2006-01-02"),
			},
			End: &calendar.EventDateTime{
				Date: am.endTime.Format("2006-01-02"),
			},
		}, nil
	}

	// if RRULE exists we cannot use all-day so we need to convert it to midnight-based time
	if am.allDay && hasRepeat {
		am.startTime = time.Date(
			am.startTime.Year(), am.startTime.Month(), am.startTime.Day(),
			0, 0, 0, 0, am.startTime.Location(),
		)
		am.endTime = am.startTime.Add(24 * time.Hour)
	}

	// parse times (for non-all-day)
	if !am.allDay {
		startParsed, err := time.Parse("15:04", am.startTimeInput.Value())
		if err != nil {
			return nil, fmt.Errorf("invalid start time format")
		}
		am.startTime = time.Date(
			am.startTime.Year(), am.startTime.Month(), am.startTime.Day(),
			startParsed.Hour(), startParsed.Minute(),
			0, 0, am.startTime.Location(),
		)

		endParsed, err := time.Parse("15:04", am.endTimeInput.Value())
		if err != nil {
			return nil, fmt.Errorf("invalid end time format")
		}
		am.endTime = time.Date(
			am.endTime.Year(), am.endTime.Month(), am.endTime.Day(),
			endParsed.Hour(), endParsed.Minute(),
			0, 0, am.endTime.Location(),
		)
	}

	tz, _ := time.Now().Local().Zone()

	// build event
	event := &calendar.Event{
		Summary:     am.title,
		Description: am.description,
		Location:    am.location,
		Start: &calendar.EventDateTime{
			DateTime: am.startTime.Format(time.RFC3339),
			TimeZone: tz,
		},
		End: &calendar.EventDateTime{
			DateTime: am.endTime.Format(time.RFC3339),
			TimeZone: tz,
		},
	}

	if hasRepeat {
		event.Recurrence = []string{rrule}
	}

	return event, nil
}

func (am *addEventModel) submitEditEventForm(event *calendar.Event) (*calendar.Event, error) {
	am.title = am.titleInput.Value()
	am.description = am.descriptionInput.Value()
	am.location = am.locationInput.Value()

	rrule := am.repeatToRRule()
	hasRepeat := rrule != ""

	// handle all-day
	if am.allDay && !hasRepeat {
		event.Summary = am.title
		event.Description = am.description
		event.Location = am.location

		event.Start = &calendar.EventDateTime{
			Date: am.startTime.Format("2006-01-02"),
		}
		event.End = &calendar.EventDateTime{
			Date: am.endTime.Format("2006-01-02"),
		}

		event.Recurrence = nil
		event.Reminders = &calendar.EventReminders{UseDefault: true}
		return event, nil
	}

	// if RRULE exists we cannot use all-day so we need to convert it to midnight-based time
	if am.allDay && hasRepeat {
		am.startTime = time.Date(
			am.startTime.Year(), am.startTime.Month(), am.startTime.Day(),
			0, 0, 0, 0, am.startTime.Location(),
		)
		am.endTime = am.startTime.Add(24 * time.Hour)
	}

	// parse time fields (if not all-day)
	if !am.allDay {
		startParsed, err := time.Parse("15:04", am.startTimeInput.Value())
		if err != nil {
			return nil, fmt.Errorf("invalid start time format")
		}
		am.startTime = time.Date(
			am.startTime.Year(), am.startTime.Month(), am.startTime.Day(),
			startParsed.Hour(), startParsed.Minute(),
			0, 0, am.startTime.Location(),
		)

		endParsed, err := time.Parse("15:04", am.endTimeInput.Value())
		if err != nil {
			return nil, fmt.Errorf("invalid end time format")
		}
		am.endTime = time.Date(
			am.endTime.Year(), am.endTime.Month(), am.endTime.Day(),
			endParsed.Hour(), endParsed.Minute(),
			0, 0, am.endTime.Location(),
		)
	}

	// update event
	event.Summary = am.title
	event.Description = am.description
	event.Location = am.location

	tz, _ := time.Now().Local().Zone()

	event.Start = &calendar.EventDateTime{
		DateTime: am.startTime.Format(time.RFC3339),
		TimeZone: tz,
	}
	event.End = &calendar.EventDateTime{
		DateTime: am.endTime.Format(time.RFC3339),
		TimeZone: tz,
	}

	if hasRepeat {
		event.Recurrence = []string{rrule}
	} else {
		event.Recurrence = nil
	}

	event.Reminders = &calendar.EventReminders{
		UseDefault: true,
	}

	return event, nil
}
