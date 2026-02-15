// Copyright (c) 2026 Romerito Mendes Silva
// Licensed under the GPLv3. See LICENSE for details.
package config

type Config struct {
	SSHPort   int
	WebPort   int
	Host      string
	AuthToken string
}

func Load() *Config {
	return &Config{
		SSHPort: 2223,
		WebPort: 8080,
		Host:    "0.0.0.0",
	}
}
