// Copyright (c) 2026 Romerito Mendes Silva
// Licensed under the GPLv3. See LICENSE for details.
package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/romeritomendes/btpterminalapp/server/internal/ui/pages"
)

type MainModel struct {
	currentPage pages.Page
	username    string
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func NewModel(username string) MainModel {
	return MainModel{
		currentPage: pages.NewHomePage(),
		username:    username,
	}
}
