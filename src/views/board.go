package views

import (
	"kiiow/gomder/config"
	"kiiow/gomder/styles"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var Board *MainView

/* MainView */
type MainView struct {
	views           []tea.Model
	ViewStatusStyle []styles.ViewStyle
	CurrentView     config.View
	infoBarView     *InfoBarView
	quitting        bool
	loaded          bool
}

/* Builder */
func NewMainView() *MainView {
	board := MainView{}
	board.initViews()
	return &board
}

/* Actions */
func (m *MainView) initViews() {
	d := NewDirectoryView()
	m.views = []tea.Model{
		d,
		NewTerminalView(d.Currentdir()),
	}
	m.ViewStatusStyle = styles.InitViewStatusStyles(len(m.views))
	m.CurrentView = config.DefaultViewIndex
}

func (m *MainView) switchView() {
	if m.CurrentView == config.TerminalIndex {
		m.CurrentView = config.DefaultViewIndex
	} else {
		m.CurrentView++
	}
}

func (m *MainView) GetViewAssociatedStyled(view config.View) lipgloss.Style {
	if view == m.CurrentView {
		return styles.FocusedStyle
	}
	return styles.BaseStyle
}

/* Rendering */
func (m *MainView) Init() tea.Cmd {
	return nil
}

func (m *MainView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.loaded {
			height, width := styles.UpdateStyleHeighAndWidth(msg.Height, msg.Width)
			m.loaded = true
			m.infoBarView = NewInfoBarView(height, width)
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			if m.CurrentView != config.TerminalIndex {
				m.quitting = true
				return m, tea.Quit
			}
		case "tab":
			m.switchView()
		}
	}
	m.infoBarView, _ = m.infoBarView.Update(msg)
	var cmd tea.Cmd
	m.views[m.CurrentView], cmd = m.views[m.CurrentView].Update(msg)
	return m, cmd
}

func (m *MainView) View() string {
	if m.quitting {
		return ""
	}
	if m.loaded {
		return styles.BoardStyle.Render(
			lipgloss.JoinVertical(
				lipgloss.Left,
				lipgloss.JoinHorizontal(
					lipgloss.Left,
					m.GetViewAssociatedStyled(config.DirectoryIndex).Render(m.views[config.DirectoryIndex].View()),
					m.GetViewAssociatedStyled(config.TerminalIndex).Render(m.views[config.TerminalIndex].View()),
				),
				m.infoBarView.View(),
			),
		)
	} else {
		return "loading ..."
	}
}
