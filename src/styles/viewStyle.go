package styles

import "kiiow/gomder/config"

type ViewStyle struct {
	Name            string
	ColorForeground string
	ColorBackground string
}

var TerminalViewStyle *ViewStyle
var DirectoryViewStyle *ViewStyle

func InitViewStatusStyles(viewLength int) []ViewStyle {
	TerminalViewStyle = &ViewStyle{
		Name:            "TERMINAL",
		ColorForeground: "0",
		ColorBackground: "36",
	}
	DirectoryViewStyle = &ViewStyle{
		Name:            "DIR",
		ColorForeground: "0",
		ColorBackground: "179",
	}
	viewsStatusStyles := make([]ViewStyle, viewLength)
	viewsStatusStyles[config.DirectoryIndex] = *DirectoryViewStyle
	viewsStatusStyles[config.TerminalIndex] = *TerminalViewStyle
	return viewsStatusStyles
}
