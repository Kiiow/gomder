package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type DirectoryView struct {
	directories table.Model
	currentdir  string
}

func CreateDirectoryView(currentdir *string) *DirectoryView {
	columns := []table.Column{
		{Title: "", Width: 3},
		{Title: "Name", Width: 15},
		{Title: "Mode", Width: 10},
	}
	entries, err := os.ReadDir(*currentdir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rows := []table.Row{}

	rows = append(rows, table.Row{
		"🔙",
		"../",
		"",
	})
	for _, entry := range entries {
		fileType := "📄"
		if entry.IsDir() {
			fileType = "📁"
		}
		rows = append(rows, table.Row{
			fileType,
			entry.Name(),
			entry.Type().Perm().String(),
		})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	d := DirectoryView{t, *currentdir}

	return &d
}

/* MAIN View */
func (d DirectoryView) Init() tea.Cmd {
	return nil
}

func (d DirectoryView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return d, tea.Quit
		}
	}
	d.directories, cmd = d.directories.Update(msg)
	return d, cmd
}

func (d DirectoryView) View() string {
	return d.directories.View()
}

func main() {
	currentdir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	d := CreateDirectoryView(&currentdir)
	p := tea.NewProgram(d)
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
