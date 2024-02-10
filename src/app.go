package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	directory_view view = iota
	terminal_view
)

const default_view = directory_view

var board *MainView

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
	if view == m.currentview {
		return focusedStyle
	}
	return baseStyle
}

/* Builder */
func (m *MainView) initViews() {
	d := NewDirectoryView()
	m.views = []tea.Model{
		d,
		NewTerminalView(d.Currentdir()),
	}
	m.currentview = default_view
}

func (m *MainView) switchView() {
	if m.currentview == terminal_view {
		m.currentview = default_view
	} else {
		m.currentview++
	}
}

func main() {
	board = NewMainView()
	board.initViews()
	p := tea.NewProgram(board)
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
