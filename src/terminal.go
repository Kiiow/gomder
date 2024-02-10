package main

import (
	"fmt"
	"path"

	tea "github.com/charmbracelet/bubbletea"
)

/* Terminal View */
type TerminalView struct {
	currentdir string
	counter    int
}

func (t *TerminalView) UpdateCurrentDir(currentdir string) tea.Msg {
	t.currentdir = path.Join(currentdir)
	return t
}

/* Builder */
func NewTerminalView(currentdir *string) *TerminalView {
	return &TerminalView{currentdir: *currentdir}
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
		case "+":
			t.counter++
		}
	}

	return t, cmd
}

func (t TerminalView) View() string {
	return fmt.Sprintf("TerminalView %s\n%v", t.currentdir, t.counter)
}
