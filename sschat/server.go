package sschat

import (
	"fmt"

	"github.com/gliderlabs/ssh"
	"github.com/patrickhener/sschat/chatuser"
	"github.com/patrickhener/sschat/logger"
	"github.com/patrickhener/sschat/rooms"
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
	logger.Infof("Starting server @ %s:%d with key %s\n", s.IP, s.Port, s.SSHKey)
	logger.Fatalf("Error starting server: %+v", ssh.ListenAndServe(addr, s.Handle(), ssh.HostKeyFile(s.SSHKey)))
}

// Handle will handle a client connecting
func (srv *Server) Handle() ssh.Handler {
	return func(s ssh.Session) {
		u := chatuser.New(s)
		if u == nil {
			return
		}
		logger.Infof("New user %s connected from %s\n", u.Nick, u.Addr)

		// Add connected user to room #general
		rooms.GeneralRoom.UsersMutex.Lock()
		rooms.GeneralRoom.Users = append(rooms.GeneralRoom.Users, u)
		rooms.GeneralRoom.UsersMutex.Unlock()

		switch len(rooms.GeneralRoom.Users) - 1 {
		case 0:
			u.Print("", "Welcome to sschat. There are no more users")
		case 1:
			u.Print("", "Welcome to sschat. There is one more user. Say hello!")
		default:
			u.Print("", fmt.Sprintf("Welcome to sschat. There are %d more users. Say hello!", len(rooms.GeneralRoom.Users)-1))
		}

		rooms.GeneralRoom.Broadcast("sschat", fmt.Sprintf("%s has joined the chat.", u.Nick))

		// Put user into chat loop for {}
		u.ChatLoop()
	}
}
