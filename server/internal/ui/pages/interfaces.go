// Copyright (c) 2026 Romerito Mendes Silva
// Licensed under the GPLv3. See LICENSE for details.
package pages

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Page interface {
	Update(msg tea.Msg) (Page, tea.Cmd)
	View() string
}

type PageModel struct {
	path    string
	content string
	error   error
}
