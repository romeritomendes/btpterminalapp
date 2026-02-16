// Copyright (c) 2026 Romerito Mendes Silva
// Licensed under the GPLv3. See LICENSE for details.
package server

import (
	"context"
	"fmt"
	"log"
	"net"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
	"github.com/romeritomendes/btpterminalapp/server/internal/config"
	"github.com/romeritomendes/btpterminalapp/server/internal/ui"
)

func StartSSH(ctx context.Context, cfg *config.Config) error {

	addr := net.JoinHostPort(cfg.Host, fmt.Sprint(cfg.SSHPort))

	server, err := wish.NewServer(
		wish.WithAddress(addr),
		wish.WithHostKeyPath(".ssh/term_info_ed25519"),
		wish.WithMiddleware(
			bubbletea.Middleware(teaHandler),
			logging.Middleware(),
		),
		wish.WithPublicKeyAuth(func(ctx ssh.Context, key ssh.PublicKey) bool { return true }),
		//wish.WithPasswordAuth(func(ctx ssh.Context, password string) bool { return true }),
	)

	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
		ctx.Done()
	}

	log.Println("Server SSH is starting...")
	log.Printf("Listening on %s\n", addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
		ctx.Done()
	}

	return nil
}

func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	m := ui.NewModel(s.User())

	return m, []tea.ProgramOption{
		tea.WithInput(s),
		tea.WithOutput(s),
	}
}
