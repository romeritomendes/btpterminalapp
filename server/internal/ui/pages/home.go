// Copyright (c) 2026 Romerito Mendes Silva
// Licensed under the GPLv3. See LICENSE for details.
package pages

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type HomeModel struct {
	choices []string
	cursor  int
}

func (m HomeModel) Update(msg tea.Msg) (Page, tea.Cmd) {
	if KeyMsg, ok := msg.(tea.KeyMsg); ok {
		switch KeyMsg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices) {
				m.cursor++
			}
		case "enter":
			switch m.choices[m.cursor] {
			case "cap":
				return &PageModel{path: "./cap_page.txt"}, nil
			case "rap":
				return &PageModel{path: "./rap_page.txt"}, nil
			case "cpi":
				return &PageModel{path: "./cpi_page.txt"}, nil
			case "bpa":
				return &PageModel{path: "./bpa_page.txt"}, nil
			case "settings":
				return &SettingsModel{}, nil
			}
		}
	}

	return m, nil
}

func (m HomeModel) View() string {
	var sb strings.Builder
	sb.WriteString("Please select an option to continue.\n\n")
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		fmt.Fprintf(&sb, "%s %s\n", cursor, choice)
	}

	return sb.String()
}

func NewHomePage() Page {
	return &HomeModel{
		choices: []string{"cap", "rap", "cpi", "bpa", "settings"},
		cursor:  0,
	}
}
