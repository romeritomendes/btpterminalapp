// Copyright (c) 2026 Romerito Mendes Silva
// Licensed under the GPLv3. See LICENSE for details.
package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/gorilla/websocket"
	"github.com/romeritomendes/btpterminalapp/server/internal/config"
	"github.com/romeritomendes/btpterminalapp/server/internal/ui"
	"golang.org/x/crypto/ssh"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandlerWS(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()

		m := ui.NewModel("ws-user")
		p := tea.NewProgram(m, tea.WithInput(conn.UnderlyingConn()), tea.WithOutput(conn.UnderlyingConn()))

		if _, err := p.Run(); err != nil {
			ctx.Done()
		}
	}
}

func HandlerProxySSH(ctx context.Context, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wsConn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Error("Failed to upgrade to WebSocket", "error", err)
			return
		}
		defer wsConn.Close()

		config := &ssh.ClientConfig{
			User:            "proxy-user",
			Auth:            []ssh.AuthMethod{ssh.Password("")},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}

		addrSSH := net.JoinHostPort("localhsot", fmt.Sprint(cfg.SSHPort))
		sshClient, err := ssh.Dial("tcp", addrSSH, config)
		if err != nil {
			log.Error("Failed to connect to SSH Server", "error", err)
			wsConn.WriteMessage(websocket.TextMessage, []byte("Failed to connect to SSH Server"))
			return
		}
		defer sshClient.Close()

		session, _ := sshClient.NewSession()
		defer session.Close()

		modes := ssh.TerminalModes{ssh.ECHO: 0}
		session.RequestPty("xterm-256color", 80, 40, modes)

		sshIn, _ := session.StdinPipe()
		sshOut, _ := session.StdoutPipe()
		session.Start("")

		log.Info("WebSocket client connected", "remote", r.RemoteAddr)

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			for {
				_, msg, err := wsConn.ReadMessage()
				if err != nil {
					log.Debug("WebSocket read error", "error", err)
					<-ctx.Done()
				}
				if _, err := sshIn.Write(msg); err != nil {
					log.Debug("SSH write error", "error", err)
					<-ctx.Done()
				}
			}
		}()

		go func() {
			defer wg.Done()
			buf := make([]byte, 32*1024)
			for {
				n, err := sshOut.Read(buf)
				if err != nil {
					log.Debug("SSH read error", "error", err)
					<-ctx.Done()
				}
				if err := wsConn.WriteMessage(websocket.BinaryMessage, buf[:n]); err != nil {
					log.Debug("WebSocket write error", "error", err)
					<-ctx.Done()
				}
			}
		}()

		wg.Wait()
		log.Info("WebSocket client disconnected", "remote", r.RemoteAddr)
	}
}
