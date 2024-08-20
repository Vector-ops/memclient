package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	NudeBeige      = "#D3B9A3"
	PalePink       = "#F4C2C2"
	SoftCoral      = "#F88379"
	Rosewood       = "#905D5D"
	BerryRed       = "#B22222"
	Mauve          = "#996666"
	PeachyNude     = "#F5CBA7"
	BlushPink      = "#FFB6C1"
	CherryRed      = "#C71585"
	Plum           = "#673147"
	ClearGloss     = "#FDFDFD"
	Champagne      = "#F7E7CE"
	ChocolateBrown = "#3B2F2F"
	DeepWine       = "#5C212A"
	WarmCaramel    = "#AF6E4D"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				tea.Printf("%s -> %s", m.table.SelectedRow()[1], m.table.SelectedRow()[2]),
			)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func Bubbletea(data map[string]string) {
	columns := []table.Column{
		{Title: "Index", Width: 20},
		{Title: "Key", Width: 100},
		{Title: "Value", Width: 80},
	}

	rows := []table.Row{}
	i := 1
	for k, v := range data {
		rows = append(rows, table.Row{
			strconv.Itoa(i),
			k,
			v,
		})
		i = i + 1
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(len(data)),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color(NudeBeige)).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color(Champagne)).
		Background(lipgloss.Color(CherryRed)).
		Bold(false)
	t.SetStyles(s)

	m := model{t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
