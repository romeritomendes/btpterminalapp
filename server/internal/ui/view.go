// Copyright (c) 2026 Romerito Mendes Silva
// Licensed under the GPLv3. See LICENSE for details.
package ui

import "fmt"

func (m MainModel) View() string {
	view := m.currentPage.View()

	s := fmt.Sprintf(
		"%s\n\nTo exit, press q or Ctrl+C.",
		view,
	)

	return s
}
