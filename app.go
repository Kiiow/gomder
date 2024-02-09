package main

import (
	"fmt"
	directory "kiiow/gomder/services/directory"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	currentdir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	d := directory.CreateDirectoryView(&currentdir)
	p := tea.NewProgram(d)
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
