package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type modelRadioModel struct {
	cursor   int
	choices  []string
	selected int
	quit     bool
}

func initialModel() *modelRadioModel {
	return &modelRadioModel{
		cursor:   0,
		choices:  []string{"微热模式", "高烧模式", "活力全开模式模式"},
		selected: 0,
		quit:     false,
	}
}

func (m *modelRadioModel) Init() tea.Cmd {
	return nil
}

func (m *modelRadioModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.quit = true
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			return m, tea.Quit
		case " ":
			m.selected = m.cursor
		}
	}

	return m, nil
}

func (m *modelRadioModel) View() string {
	s := "选择你的暖手温度?\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if i == m.selected {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += `
1. 按空格选择或取消选择
2. 按上下键切换候选
3. 按 ctrl+c 退出
`
	return s
}
