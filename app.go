package main

import (
	"fmt"
	"os"
	"path"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type DirectoryView struct {
	directories table.Model
	currentdir  string
	quitting    bool
}

type columntype int

const (
	directory_filetype columntype = iota
	directory_filename
	directory_filemode
)

const (
	emoji_back   string = "🔙"
	emoji_folder string = "📁"
	emoji_file   string = "📄"
)

func CreateDirectoryView(currentdir *string) *DirectoryView {
	columns := []table.Column{
		{Title: "", Width: 3},
		{Title: "Name", Width: 15},
		{Title: "Mode", Width: 10},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithFocused(true),
		table.WithHeight(15),
	)

	d := DirectoryView{t, *currentdir, false}
	d.directories.SetRows(*d.UpdateDirectory())

	return &d
}

func (d *DirectoryView) UpdateDirectory() *[]table.Row {
	entries, err := os.ReadDir(d.currentdir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rows := []table.Row{}

	rows = append(rows, table.Row{
		emoji_back,
		"../",
		"",
	})
	for _, entry := range entries {
		fileType := emoji_file
		if entry.IsDir() {
			fileType = emoji_folder
		}
		rows = append(rows, table.Row{
			fileType,
			entry.Name(),
			entry.Type().Perm().String(),
		})
	}
	return &rows
}

func (d *DirectoryView) MoveIn(folder string) {
	d.currentdir = path.Join(d.currentdir, folder)
}

func (d *DirectoryView) MoveOut() {
	d.currentdir = path.Join(d.currentdir, "..")
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
			d.quitting = true
			return d, tea.Quit
		case "enter", "space":
			row := d.directories.SelectedRow()
			c := row[directory_filetype]
			if c == emoji_folder || c == emoji_back {
				if d.directories.Cursor() == 0 {
					d.MoveOut()
				} else {
					d.MoveIn(row[directory_filename])
				}
				d.directories.SetRows(*d.UpdateDirectory())
				d.directories.GotoTop()
			}
		}
	}
	d.directories, cmd = d.directories.Update(msg)
	return d, cmd
}

func (d DirectoryView) View() string {
	// Avoid quitting app and rerendering
	if d.quitting {
		return ""
	}
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
