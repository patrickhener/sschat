package sschat

import (
	"fmt"

	"github.com/gliderlabs/ssh"
	"github.com/patrickhener/sschat/chatuser"
)

// Server will represent the ssh server
type Server struct {
	Port   int
	IP     string
	SSHKey string
}

// Run will start the ssh server
func (s *Server) Run() {
	addr := fmt.Sprintf("%s:%d", s.IP, s.Port)
	fmt.Printf("Starting server @ %s:%d with key %s\n", s.IP, s.Port, s.SSHKey)
	panic(ssh.ListenAndServe(addr, s.Handle(), ssh.HostKeyFile(s.SSHKey)))
}

// Handle will handle a client connecting
func (srv *Server) Handle() ssh.Handler {
	return func(s ssh.Session) {
		fmt.Printf("New user connected from %s\n", s.RemoteAddr())
		u := chatuser.New(s)
		fmt.Println(u)
	}
}
