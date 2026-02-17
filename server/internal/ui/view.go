// Copyright (c) 2026 Romerito Mendes Silva
// Licensed under the GPLv3. See LICENSE for details.
package ui

import (
	"fmt"
	//"github.com/charmbracelet/log"
)

func (m MainModel) View() string {
	if m.currentPage == nil {
		//log.Info("Mission accomplished. Goodbye, human!")
		return "\n\nMission accomplished. Goodbye, human!"
	}

	view := m.currentPage.View()

	s := fmt.Sprintf(
		"%s\n\nTo exit, press q or Ctrl+C.",
		view,
	)

	return s
}
