package models

import (
	"github.com/NStefan002/tui-calendar/v2/styles"
	"github.com/NStefan002/tui-calendar/v2/utils"

	"github.com/charmbracelet/lipgloss"
)

func errorView(errMsg string, scrWidth int) string {
	content := lipgloss.JoinVertical(
		lipgloss.Left,
		styles.ErrorTitleStyle.Render("⚠ An error occurred"),
		"",
		errMsg,
		styles.ErrorHintStyle.Render("Press any key to continue"),
	)

	box := styles.ErrorBoxStyle.Render(content)

	return utils.CenterText(box, scrWidth)
}
