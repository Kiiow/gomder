package styles

import "github.com/charmbracelet/lipgloss"

/* Styling */
var (
	BoardStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("15"))
	/* Components border */
	BaseStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("240"))
	FocusedStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62"))
)

func UpdateStyleHeighAndWidth(height, width int) {
	heightFixed := height - 2
	widthFixed := width - 2

	BoardStyle.Height(heightFixed)
	BoardStyle.Width(widthFixed)

	BaseStyle.Height(heightFixed - 4)
	FocusedStyle.Height(heightFixed - 4)
}
