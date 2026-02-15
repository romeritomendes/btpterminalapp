// Copyright (c) 2026 Romerito Mendes Silva
// Licensed under the GPLv3. See LICENSE for details.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/romeritomendes/btpterminalapp/server/internal/config"
	"github.com/romeritomendes/btpterminalapp/server/internal/server"
	"github.com/romeritomendes/btpterminalapp/server/internal/ui"
)

func main() {
	if os.Getenv("profile") != "local" {
		cfg := config.Load()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		errChan := make(chan error, 2)

		//SSH Server
		go func() {
			errChan <- server.StartSSH(ctx, cfg)
		}()

		//Web Server
		go func() {

		}()

		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

		select {
		case err := <-errChan:
			log.Fatalf("Internal service failure: %v", err)
			cancel()

		case sig := <-sigChan:
			log.Printf("Operation cancelled by user request. %v", sig)
			cancel()
		}
	} else {
		p := tea.NewProgram(ui.NewModel("localUser"))
		if _, err := p.Run(); err != nil {
			fmt.Printf("Erro %v", err)
			os.Exit(1)
		}
	}
}
