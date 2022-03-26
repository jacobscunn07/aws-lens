package ui

import (
  "fmt"
  "github.com/charmbracelet/lipgloss"
)

type HeaderPanelData struct {
  Account   string
  Region    string
  Principal string
}

type HeaderPanelView struct {
  //style lipgloss.Style
  data HeaderPanelData
}

func NewHeaderPanelView(d HeaderPanelData) *HeaderPanelView {
  return &HeaderPanelView{
    //style: lipgloss.Style{},
    data: d,
  }
}

func (v *HeaderPanelView) Render() string {
  return lipgloss.JoinHorizontal(
    lipgloss.Top,
    lipgloss.JoinVertical(
      lipgloss.Top,
      "Principal: ",
      "Region: ",
      "Account: ",
      "Profile: ",
    ),
    lipgloss.JoinVertical(
      lipgloss.Top,
      fmt.Sprintf(v.data.Principal),
      fmt.Sprintf(v.data.Region),
      fmt.Sprintf(v.data.Account),
      "default",
    ),
  )
}
