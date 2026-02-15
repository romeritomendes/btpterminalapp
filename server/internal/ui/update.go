// Copyright (c) 2026 Romerito Mendes Silva
// Licensed under the GPLv3. See LICENSE for details.
package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/romeritomendes/btpterminalapp/server/internal/ui/pages"
)

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "esc":
			m.currentPage = pages.NewHomePage()
		}

	case pages.Page:
		m.currentPage = msg
	default:
		m.currentPage = pages.NewHomePage()
	}

	var cmd tea.Cmd
	m.currentPage, cmd = m.currentPage.Update(msg)
	return m, cmd
}
