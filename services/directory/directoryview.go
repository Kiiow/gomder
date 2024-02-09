package directory

import (
	"fmt"
	"kiiow/gomder/services/format"
	"os"
	"path"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

/* Directory View */
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

/* Builder */
func NewView() *DirectoryView {
	currentdir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	columns := []table.Column{
		{Title: "", Width: 2},
		{Title: "Name", Width: 15},
		{Title: "Size", Width: 9},
		{Title: "Mode", Width: 10},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithFocused(true),
		table.WithHeight(15),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("147")).
		Bold(true)

	t.SetStyles(s)

	d := DirectoryView{t, currentdir, false}
	d.directories.SetRows(*d.UpdateDirectory())

	return &d
}

/* Actions */
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
		"",
	})
	for _, entry := range entries {
		fileType := emoji_file
		if entry.IsDir() {
			fileType = emoji_folder
		}
		i, err := entry.Info()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		perms := i.Mode() & os.ModePerm
		rows = append(rows, table.Row{
			fileType,
			entry.Name(),
			fmt.Sprintf("%9v", format.Bytes_fmt(i.Size())),
			fmt.Sprintf("%v", perms),
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

func (d *DirectoryView) Move() {
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

/* Rendering */
func (d DirectoryView) Init() tea.Cmd {
	return nil
}

func (d DirectoryView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			d.Move()
		}
	}
	d.directories, cmd = d.directories.Update(msg)
	return d, cmd
}

func (d DirectoryView) View() string {
	if d.quitting {
		return ""
	}
	tableeol := ""
	if len(d.directories.Rows()) > d.directories.Height() {
		tableeol = " ..."
	}
	return d.directories.View() + fmt.Sprintf("\n%s", tableeol)
}
