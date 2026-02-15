// Copyright (c) 2026 Romerito Mendes Silva
// Licensed under the GPLv3. See LICENSE for details.
package pages

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type fileLoadedMsg string
type errMsg error

func readFile(path string) tea.Cmd {
	return func() tea.Msg {
		content, err := os.ReadFile(path)
		if err != nil {
			return errMsg(err)
		}

		return fileLoadedMsg(string(content))
	}
}

func (m PageModel) Update(msg tea.Msg) (Page, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "r" {
			return m, readFile(m.path)
		}
	case fileLoadedMsg:
		m.content = string(msg)
		m.error = nil
		return m, nil
	case errMsg:
		m.error = msg
	}

	return m, nil
}

func (m PageModel) View() string {
	var s string = "Press esc to return to options\n\n"

	if m.error != nil {
		return fmt.Sprintf("%sFailed to read the file: %s", s, m.error)
	}
	if m.content == "" {
		return fmt.Sprintf("%sLoading file, please wait...", s)
	}
	return fmt.Sprintf("%s%s", s, m.content)
}
