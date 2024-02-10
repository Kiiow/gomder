package views

import (
	"fmt"
	"kiiow/gomder/format"
	"os"
	"path/filepath"
	"sort"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type DirectoryPath struct {
	path string
}

func (d *DirectoryPath) Path() string {
	return d.path
}

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
	emoji_bipper string = "📟"
)

func (d *DirectoryView) Currentdir() *string {
	return &d.currentdir
}

/* Styling */
func GetStyling() table.Styles {
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

	return s
}

/* Builder */
func NewDirectoryView() *DirectoryView {
	currentdir, err := filepath.Abs(".")
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

	t.SetStyles(GetStyling())

	d := DirectoryView{
		directories: t,
		currentdir:  currentdir,
		quitting:    false,
	}
	d.directories.SetRows(*d.updateDirectory())

	return &d
}

/* Actions */
func (d *DirectoryView) updateDirectory() *[]table.Row {
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

	// Sort folder first and in alphabetic order
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].IsDir() == entries[j].IsDir() {
			return entries[i].Name() < entries[j].Name()
		}
		return true
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

func (d *DirectoryView) moveIn(folder string) {
	var err error
	d.currentdir, err = filepath.Abs(filepath.Join(d.currentdir, folder))
	if err != nil {
		os.Exit(1)
	}
}

func (d *DirectoryView) moveOut() {
	var err error
	d.currentdir, err = filepath.Abs(filepath.Join(d.currentdir, ".."))
	if err != nil {
		os.Exit(1)
	}
}

func (d *DirectoryView) move() tea.Msg {
	row := d.directories.SelectedRow()
	c := row[directory_filetype]
	if c == emoji_back || c == emoji_folder {
		if d.directories.Cursor() == 0 {
			d.moveOut()

		} else {
			d.moveIn(row[directory_filename])
		}
		d.directories.SetRows(*d.updateDirectory())
		d.directories.GotoTop()
	}

	return &DirectoryPath{path: d.currentdir}
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
			d.move()
			Board.Update(d)
			return d, cmd
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
