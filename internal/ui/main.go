package ui

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var (
	physicalWidth, _, _ = term.GetSize(int(os.Stdout.Fd()))
	highlight           = lipgloss.AdaptiveColor{Light: "#000000", Dark: "#FF9900"}
)

type Component interface {
	Render() string
}

// Interfaces

type PanelView interface {
	Render() string
}

type PanelData interface {
}

//type PanelData interface{}

type SidePanelView interface {
	SidePanelViewName() string
	PanelView
}

type SidePanelData interface {
	SidePanelDataName() string
	PanelData
}

//type SidePanelData interface {
//	PanelData
//}

type DetailPanelView interface {
	DetailPanelViewName() string
	PanelView
}

// Home Side Panel View

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

// Home Detail Panel View

type HomeDetailPanelView struct {
}

func (v *HomeDetailPanelView) DetailPanelViewName() string {
	return "home-detailpanel"
}

func (v *HomeDetailPanelView) Render() string {
	return ""
}

// Header Panel View

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

type ViewModel struct {
	HeaderPanelData HeaderPanelData
	SidePanelData   PanelData
}

func (m *ViewModel) Init() tea.Cmd {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		//log.Fatalf("unable to load SDK config, %v", err)
	}

	client := sts.NewFromConfig(cfg)

	out, _ := client.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})

	m.HeaderPanelData = HeaderPanelData{
		Account:   *out.Account,
		Region:    cfg.Region,
		Principal: *out.Arn,
	}

	m.SidePanelData = HomeSidePanelData{
		Services: []string{
			"VPC",
			"EC2",
			"Systems Manager",
			"Route53",
			"RDS",
			"Lambda",
		},
	}

	return nil
}

func (m *ViewModel) View() string {
	doc := strings.Builder{}
	docStyle := lipgloss.NewStyle().Padding(1, 2, 1, 2)

	doc.WriteString(NewShell(m).Render())

	return docStyle.Render(doc.String())
}

func (m *ViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, key.NewBinding(key.WithKeys("ctrl+c", "q"))):
			return m, tea.Quit
		}
	}

	return m, nil
}

func New() error {
	m := ViewModel{}

	return tea.NewProgram(&m).Start()
}
