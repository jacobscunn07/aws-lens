package ui

import "github.com/charmbracelet/lipgloss"

type Shell struct {
  style lipgloss.Style
  model ViewModel
}

func NewShell(vm *ViewModel) *Shell {
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
    Width(physicalWidth - 4)

  return &Shell{
    model: *vm,
    style: containerStyle,
  }
}

func (s *Shell) Render() string {
  var sidePanelView SidePanelView
  switch d := s.model.SidePanelData.(type) {
  case HomeSidePanelData:
    sidePanelView = NewHomeSidePanelView(d)
    //doc.WriteString(v.Render())
  }

  return s.style.Render(lipgloss.JoinVertical(lipgloss.Top,
    NewHeaderPanelView(s.model.HeaderPanelData).Render(),
    lipgloss.JoinHorizontal(
      lipgloss.Top,
      sidePanelView.Render(),
      `
         ▄              ▄
        ▌▒█           ▄▀▒▌
        ▌▒▒█        ▄▀▒▒▒▐
       ▐▄█▒▒▀▀▀▀▄▄▄▀▒▒▒▒▒▐
     ▄▄▀▒▒▒▒▒▒▒▒▒▒▒█▒▒▄█▒▐
   ▄▀▒▒▒░░░▒▒▒░░░▒▒▒▀██▀▒▌
  ▐▒▒▒▄▄▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▀▄▒▌
  ▌░░▌█▀▒▒▒▒▒▄▀█▄▒▒▒▒▒▒▒█▒▐
 ▐░░░▒▒▒▒▒▒▒▒▌██▀▒▒░░░▒▒▒▀▄▌
 ▌░▒▒▒▒▒▒▒▒▒▒▒▒▒▒░░░░░░▒▒▒▒▌
▌▒▒▒▄██▄▒▒▒▒▒▒▒▒░░░░░░░░▒▒▒▐
▐▒▒▐▄█▄█▌▒▒▒▒▒▒▒▒▒▒░▒░▒░▒▒▒▒▌
▐▒▒▐▀▐▀▒▒▒▒▒▒▒▒▒▒▒▒▒░▒░▒░▒▒▐
 ▌▒▒▀▄▄▄▄▄▄▒▒▒▒▒▒▒▒░▒░▒░▒▒▒▌
 ▐▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒░▒░▒▒▄▒▒▐
  ▀▄▒▒▒▒▒▒▒▒▒▒▒▒▒░▒░▒▄▒▒▒▒▌
    ▀▄▒▒▒▒▒▒▒▒▒▒▄▄▄▀▒▒▒▒▄▀
      ▀▄▄▄▄▄▄▀▀▀▒▒▒▒▒▄▄▀
         ▀▀▀▀▀▀▀▀▀▀▀▀
`,
    )))
}
