package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type view int

/* MainView */
type MainView struct {
	views       []tea.Model
	currentview view
	quitting    bool
	loaded      bool
}

/* Builder */
func NewMainView() *MainView {
	return &MainView{}
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
	case DirectoryView:
		m.views[terminal_view], _ = m.views[terminal_view].Update(msg)
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "tab":
			m.switchView()
		}
	}
	var cmd tea.Cmd
	m.views[m.currentview], cmd = m.views[m.currentview].Update(msg)
	return m, cmd
}

func (m MainView) View() string {
	if m.quitting {
		return ""
	}
	if m.loaded {
		return baseStyle.Render(
			lipgloss.JoinHorizontal(
				lipgloss.Left,
				m.getViewStyled(directory_view).Render(m.views[directory_view].View()),
				m.getViewStyled(terminal_view).Render(m.views[terminal_view].View()),
			),
		)
	} else {
		return "loading ..."
	}
}
