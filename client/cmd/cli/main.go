// Copyright (c) 2026 Romerito Mendes Silva
// Licensed under the GPLv3. See LICENSE for details.
package main

import (
	"context"

	"github.com/romeritomendes/btpterminalapp/client/internal/bridge"
	"github.com/romeritomendes/btpterminalapp/client/internal/config"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.Load()

	bridge.ConnectSSH(ctx, cfg)
}
