package styles

import "github.com/charmbracelet/lipgloss"

// var style = lipgloss.NewStyle().Bold(true)

var InitStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("12"))

var InitError = lipgloss.NewStyle().
	Foreground(lipgloss.Color("1"))

var InitSuccess = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#30da00"))

var AddFileStlyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#b600da"))

var Warning = lipgloss.NewStyle().Foreground(lipgloss.Color("13"))

var PromptStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#00BDFF"))

var StatusSuccess = lipgloss.NewStyle().Bold(true).
	Foreground(lipgloss.Color("#74FF33"))

var StatusPrompt = lipgloss.NewStyle().Bold(true).
	Foreground(lipgloss.Color("#00BDFF"))
