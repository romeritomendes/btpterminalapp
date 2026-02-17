// Copyright (c) 2026 Romerito Mendes Silva
// Licensed under the GPLv3. See LICENSE for details.
package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/romeritomendes/btpterminalapp/server/internal/config"
)

func StartWeb(ctx context.Context, cfg *config.Config) error {
	addr := net.JoinHostPort(cfg.Host, fmt.Sprint(cfg.WebPort))

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("public")))
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	mux.HandleFunc("/ws", HandlerWS(ctx))
	mux.HandleFunc("/proxySSH", HandlerProxySSH(ctx, cfg))

	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		srv.Shutdown(context.Background())
	}()

	log.Info("Starting Web Server", "addr", addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Error("Error to start the server", "error", err)
		return err
	}

	return nil
}
