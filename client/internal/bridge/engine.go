// Copyright (c) 2026 Romerito Mendes Silva
// Licensed under the GPLv3. See LICENSE for details.
package bridge

import (
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/romeritomendes/btpterminalapp/client/internal/config"
	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

func ConnectSSH(ctx context.Context, cfg *config.Config) {
	key, err := os.ReadFile(cfg.KeyPath)
	if err != nil {
		panic(err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		panic(err)
	}

	config := &ssh.ClientConfig{
		User: cfg.User,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         10 * time.Second,
	}

	conn, err := net.DialTimeout("tcp", cfg.Target, 10*time.Second)
	if err != nil {
		panic(err)
	}

	sshConn, chans, reqs, err := ssh.NewClientConn(conn, cfg.Target, config)
	if err != nil {
		panic(err)
	}

	client := ssh.NewClient(sshConn, chans, reqs)
	defer client.Close()

	sess, err := client.NewSession()
	if err != nil {
		panic(err)
	}
	defer sess.Close()

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	width, height, _ := term.GetSize(int(os.Stdout.Fd()))
	err = sess.RequestPty("xterm-256color", height, width, ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	})
	if err != nil {
		panic(err)
	}

	sess.Stdout = os.Stdout
	sess.Stderr = os.Stderr
	sess.Stdin = os.Stdin

	if err := sess.Shell(); err != nil {
		panic(err)
	}

	go func() {
		prevW, prevH := width, height
		for {
			w, h, err := term.GetSize(int(os.Stdout.Fd()))
			if err == nil && w > 0 && h > 0 && (w != prevW || h != prevH) {
				_ = sess.WindowChange(h, w)
				prevW, prevH = w, h
			}
			time.Sleep(300 * time.Millisecond)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		sess.Signal(ssh.SIGINT)
	}()

	err = sess.Wait()
	if err != nil && err != io.EOF {
		fmt.Println("SSH closed with error:", err)
	}
}
