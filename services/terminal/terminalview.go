package terminal

import (
	tea "github.com/charmbracelet/bubbletea"
)

/* Terminal View */
type TerminalView struct {
}

/* Builder */
func NewView() *TerminalView {
	return &TerminalView{}
}

func (t TerminalView) Init() tea.Cmd {
	return nil
}

func (t TerminalView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// Do stuff in here
		}
	}

	return t, cmd
}

func (t TerminalView) View() string {
	return "TerminalView"
}
