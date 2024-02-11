package views

import (
	"fmt"
	"os/exec"
	"strings"

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
				Foreground(lipgloss.Color("111"))
)

func ExecuteCommand(dir string, command string) (string, error) {
	cmdArgs := strings.Split(command, " ")

	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Dir = dir

	res, err := cmd.Output()
	var output string
	if err != nil {
		return output, err
	}
	output = fmt.Sprintf("%s", res)
	return output, nil
}

/* Builder */
func NewTerminalView(currentdir *string, height, width int) *TerminalView {
	ti := textinput.New()
	ti.Prompt = input_default_prompt
	ti.Focus()
	ti.Width = 50

	return &TerminalView{
		currentdir:   *currentdir,
		commandInput: ti,
	}
}

func (t *TerminalView) Init() tea.Cmd {
	return textinput.Blink
}

func (t *TerminalView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			t.commandOutput, _ = ExecuteCommand(t.currentdir, t.lastExecutedCommand)
			t.commandInput.SetValue("")
		}
	}
	t.commandInput, cmd = t.commandInput.Update(msg)

	return t, cmd
}

func (t *TerminalView) View() string {
	return workingdirectory.Render(t.currentdir) +
		fmt.Sprintf("\n%s", t.lastExecutedCommand) +
		t.commandInput.View() +
		fmt.Sprintf("\n%s", t.commandOutput)
}
