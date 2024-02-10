package config

type View int

const (
	DirectoryIndex View = iota
	TerminalIndex
)

const (
	DirectoryWidthPercent = 30
	TerminalWidthPercent  = 70
)

const DefaultViewIndex = DirectoryIndex
