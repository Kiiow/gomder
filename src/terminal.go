package main

import (
	"fmt"
	"os/exec"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

/* Terminal View */
type TerminalView struct {
	currentdir          string
	lastExecutedCommand string
	commandOutput       string
	commandInput        textinput.Model
}

func (t *TerminalView) UpdateCurrentDir(moveTo string) tea.Msg {
	t.currentdir = moveTo
	return t
}

/* Styling */
var (
	input_default_prompt = fmt.Sprintf("\n%s > ", emoji_bipper)
	workingdirectory     = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("63"))
)

/* Builder */
func NewTerminalView(currentdir *string) *TerminalView {
	ti := textinput.New()
	ti.Prompt = input_default_prompt
	ti.Focus()
	ti.Width = 50

	return &TerminalView{currentdir: *currentdir, commandInput: ti}
}

func (t TerminalView) Init() tea.Cmd {
	return textinput.Blink
}

func (t TerminalView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case DirectoryView:
		t.UpdateCurrentDir(msg.currentdir)
		return t, cmd
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			t.commandInput.SetValue("")
		case "enter":
			t.lastExecutedCommand = t.commandInput.Value()
			cmd := exec.Command(t.lastExecutedCommand)
			cmd.Dir = t.currentdir
			res, err := cmd.Output()
			if err != nil {
				t.commandOutput = err.Error()
			} else {
				t.commandOutput = fmt.Sprint(res)
			}
			t.commandInput.SetValue("")
		}
	}
	t.commandInput, cmd = t.commandInput.Update(msg)

	return t, cmd
}

func (t TerminalView) View() string {
	return workingdirectory.Render(t.currentdir) +
		fmt.Sprintf("\n%s", t.lastExecutedCommand) +
		t.commandInput.View() +
		fmt.Sprintf("\n%s", t.commandOutput)
}
