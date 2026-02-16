// Copyright (c) 2026 Romerito Mendes Silva
// Licensed under the GPLv3. See LICENSE for details.
package server

import (
	"context"
	"net/http"

	"github.com/romeritomendes/btpterminalapp/server/internal/config"
)

func StartWeb(ctx context.Context, cfg *config.Config) error {
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("public")))

	err := http.ListenAndServe(":8080", mux)
	return err
}
