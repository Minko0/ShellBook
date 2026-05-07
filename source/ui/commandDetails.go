package ui

import (
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type DetailsModel struct {
	list    list.Model
	options []option
}

type option struct {
	title       string
	description string
}

func (i option) Title() string       { return i.title }
func (i option) Description() string { return i.description }
func (i option) FilterValue() string { return i.title }

func CreateDetails() DetailsModel {
	optionsList := list.New([]list.Item{}, list.NewDefaultDelegate(), 80, 40)
	return DetailsModel{optionsList, []option{}}
}

func (details *DetailsModel) SetOptions(options []option) {
	details.options = options
	items := make([]list.Item, 0)
	for _, o := range options {
		items = append(items, o)
	}
	details.list.SetItems(items)
}

func (details *DetailsModel) UpdateDetails(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	details.list, cmd = details.list.Update(msg)
	return cmd
}

func (details *DetailsModel) View(width int, height int) string {
	termWidth := width
	topHeight := height - 2

	leftWidth := termWidth * 50 / 100
	rightWidth := termWidth - leftWidth

	left := lipgloss.NewStyle().
		Width(leftWidth).
		Height(topHeight).
		Render("Option")

	right := lipgloss.NewStyle().
		Width(rightWidth).
		Height(topHeight).
		Render(details.list.View())

	return lipgloss.JoinHorizontal(lipgloss.Top, left, right)
}
