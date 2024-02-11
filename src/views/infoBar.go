package views

import (
	"fmt"
	"kiiow/gomder/config"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type InfoBarView struct {
	maxHeight int
	maxWidth  int
	counter   int
}

/* Styling */
var (
	defaultInfoBarStyle = lipgloss.NewStyle().
				Background(lipgloss.Color("240"))
	defaultStatusActiveViewStyle = lipgloss.NewStyle().
					Foreground(lipgloss.Color("14")).
					MarginLeft(1)
)

/* Builder */
func NewInfoBarView(height, width int) *InfoBarView {
	return &InfoBarView{
		maxHeight: height,
		maxWidth:  width,
	}
}

/* Rendering */
func (i InfoBarView) Init() tea.Cmd {
	return nil
}

func (i InfoBarView) Update(msg tea.Msg) (*InfoBarView, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "h":
			i.counter++
		}
	}
	return &i, nil
}

func (i InfoBarView) View() string {
	currentViewStyle := Board.ViewStatusStyle[Board.CurrentView]

	statusStyle := defaultStatusActiveViewStyle.
		Background(lipgloss.Color(currentViewStyle.ColorBackground)).
		Foreground(lipgloss.Color(currentViewStyle.ColorForeground))
	statusText := fmt.Sprintf(" %s ", currentViewStyle.Name)
	infoBarStyle := defaultInfoBarStyle.Width(i.maxWidth - len(statusText))

	return statusStyle.Render(statusText) +
		infoBarStyle.Render(fmt.Sprintf(" %s ", config.Config.Application.Name))
}
