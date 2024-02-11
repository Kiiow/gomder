package main

import (
	"fmt"
	"kiiow/gomder/config"
	"kiiow/gomder/views"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	config.Load()
	views.Board = views.NewMainView()
	p := tea.NewProgram(views.Board)
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
