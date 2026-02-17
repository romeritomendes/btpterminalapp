// Copyright (c) 2026 Romerito Mendes Silva
// Licensed under the GPLv3. See LICENSE for details.
package config

type Config struct {
	Target  string
	User    string
	KeyPath string
}

func Load() *Config {
	return &Config{
		Target:  "localhost:2223",
		User:    "local",
		KeyPath: ".ssh/term_info_ed25519",
	}
}
