package ui

import (
	"fmt"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

type state int

const (
	listView state = iota
	detailView
)

type model struct {
	state           state
	list            list.Model
	initialCommands []string
	cursor          int
	selected        map[int]struct{}
	selectedCommand string
}

type Tui struct {
	model
}

func New(commands []string) *Tui {
	return &Tui{
		model: initialModel(commands),
	}
}

func (t *Tui) Run() {
	p := tea.NewProgram(t)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
	}
}

func initialModel(commands []string) model {
	return model{
		list:            CreateListModel(commands),
		initialCommands: commands,
		selected:        make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, msg.Height)
		return m, nil
	}

	switch m.state {
	case listView:
		return UpdateList(msg, m)

	case detailView:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "q", "ctrl+c":
				return m, tea.Quit
			case "esc", "backspace":
				m.state = listView
				return m, nil
			}
		}
	}

	return m, nil
}

func (m model) View() tea.View {
	switch m.state {
	case listView:
		return tea.NewView(m.list.View())
	case detailView:
		return tea.NewView(m.selectedCommand)
	default:
		panic("invalid state")
	}
}
