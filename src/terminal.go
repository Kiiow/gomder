package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

/* Terminal View */
type TerminalView struct {
	currentdir string
}

func (t *TerminalView) UpdateCurrentDir(moveTo string) tea.Msg {
	t.currentdir = moveTo
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
	case DirectoryView:
		t.UpdateCurrentDir(msg.currentdir)
		return t, cmd
	case tea.KeyMsg:
		switch msg.String() {
		// Do stuff in here
		}
	}

	return t, cmd
}

func (t TerminalView) View() string {
	return fmt.Sprintf("TerminalView\n%s", t.currentdir)
}
