package models

import (
	"github.com/NStefan002/tui-calendar/v2/styles"
	"github.com/NStefan002/tui-calendar/v2/utils"

	"github.com/charmbracelet/lipgloss"
)

func errorView(m *model) string {
content := lipgloss.JoinVertical(
		lipgloss.Left,
		styles.ErrorTitleStyle.Render("⚠ An error occurred"),
		"",
		m.errMessage,
		styles.ErrorHintStyle.Render("Press any key to continue"),
	)

	box := styles.ErrorBoxStyle.Render(content)

	return utils.CenterText(box, m.screenWidth)
}
