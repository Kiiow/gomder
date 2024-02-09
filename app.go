package main

import (
	"fmt"
	directory "kiiow/gomder/services/directory"
	"kiiow/gomder/services/terminal"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type view int

const (
	directory_view view = iota
	terminal_view
)
const default_view = directory_view

type MainView struct {
	views    []tea.Model
	focused  view
	quitting bool
	loaded   bool
}

/* Styling */
var (
	baseStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("240"))

	focusedStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62"))
)

func (m *MainView) getViewStyled(view view) lipgloss.Style {
	if view == m.focused {
		return focusedStyle
	}
	return baseStyle
}

/* Builder */
func New() *MainView {
	return &MainView{}
}

func (m *MainView) initViews(width, height int) {
	m.views = []tea.Model{
		directory.NewView(),
		terminal.NewView(),
	}
	m.focused = default_view
}

func (m *MainView) switchView() {
	if m.focused == terminal_view {
		m.focused = default_view
	} else {
		m.focused++
	}
}

func (m MainView) Init() tea.Cmd {
	return nil
}

func (m MainView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.loaded {
			m.initViews(msg.Width, msg.Height)
			m.loaded = true
		}
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
	m.views[m.focused], cmd = m.views[m.focused].Update(msg)
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

func main() {
	m := New()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
