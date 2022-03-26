package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
	"os"
)

var (
	physicalWidth, _, _ = term.GetSize(int(os.Stdout.Fd()))
	highlight           = lipgloss.AdaptiveColor{Light: "#000000", Dark: "#FF9900"}
)

func New() error {
	m := ViewModel{}

	return tea.NewProgram(&m).Start()
}
