// Copyright (c) 2026 Romerito Mendes Silva
// Licensed under the GPLv3. See LICENSE for details.
package pages

import tea "github.com/charmbracelet/bubbletea"

type SettingsModel struct {
	choises []string
	cursor  int
}

func (m SettingsModel) Update(tea.Msg) (Page, tea.Cmd) {
	return m, nil
}

func (m SettingsModel) View() string {
	return "Settings"
}
