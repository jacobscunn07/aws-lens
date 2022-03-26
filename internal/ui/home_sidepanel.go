package ui

import "github.com/charmbracelet/lipgloss"

type HomeSidePanelData struct {
  Services []string
}

type HomeSidePanelView struct {
  style lipgloss.Style
  data  HomeSidePanelData
}

func NewHomeSidePanelView(d HomeSidePanelData) *HomeSidePanelView {
  containerBorder := lipgloss.Border{
    Top:         "─",
    Bottom:      "─",
    Left:        "│",
    Right:       "│",
    TopLeft:     "╭",
    TopRight:    "╮",
    BottomLeft:  "╰",
    BottomRight: "╯",
  }

  containerStyle := lipgloss.NewStyle().
    Border(containerBorder, true).
    BorderForeground(highlight).
    Padding(0, 1).
    Width(physicalWidth / 100 * 20)

  return &HomeSidePanelView{
    data:  d,
    style: containerStyle,
  }
}

func (v *HomeSidePanelView) SidePanelViewName() string {
  return "home-sidepanel"
}

func (v *HomeSidePanelView) Render() string {
  return v.style.Render(lipgloss.JoinVertical(
    lipgloss.Top,
    v.data.Services...))
}
