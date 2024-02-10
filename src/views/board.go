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
	views       []tea.Model
	currentView config.View
	quitting    bool
	loaded      bool
}

func (m *MainView) CurrentView() config.View {
	return m.currentView
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
	m.currentView = config.DefaultViewIndex
}

func (m *MainView) switchView() {
	if m.currentView == config.TerminalIndex {
		m.currentView = config.DefaultViewIndex
	} else {
		m.currentView++
	}
}

func (m *MainView) GetViewAssociatedStyled(view config.View) lipgloss.Style {
	if view == m.currentView {
		return styles.FocusedStyle
	}
	return styles.BaseStyle
}

/* Rendering */
func (m MainView) Init() tea.Cmd {
	return nil
}

func (m MainView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.loaded {
			m.loaded = true
		}
		styles.UpdateStyleHeighAndWidth(msg.Height, msg.Width)
	case DirectoryView:
		m.views[config.TerminalIndex], _ = m.views[config.TerminalIndex].Update(msg)
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			if m.currentView != config.TerminalIndex {
				m.quitting = true
				return m, tea.Quit
			}
		case "tab":
			m.switchView()
		}
	}
	var cmd tea.Cmd
	m.views[m.currentView], cmd = m.views[m.currentView].Update(msg)
	return m, cmd
}

func (m MainView) View() string {
	if m.quitting {
		return ""
	}
	if m.loaded {
		return styles.BoardStyle.Render(
			lipgloss.JoinHorizontal(
				lipgloss.Left,
				m.GetViewAssociatedStyled(config.DirectoryIndex).Render(m.views[config.DirectoryIndex].View()),
				m.GetViewAssociatedStyled(config.TerminalIndex).Render(m.views[config.TerminalIndex].View()),
			),
		)
	} else {
		return "loading ..."
	}
}
