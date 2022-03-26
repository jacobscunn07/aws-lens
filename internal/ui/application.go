package ui

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

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
