package chatuser

import (
	"fmt"
	"io"
	"net"
	"strings"
	"sync"
	"time"

	terminal "golang.org/x/term"

	"github.com/gliderlabs/ssh"
	"github.com/patrickhener/sschat/logger"
	"github.com/patrickhener/sschat/utils"
)

var (
	// Init empty users map
	AllUsers      = make(map[string]string, 400)
	AllUsersMutex = sync.Mutex{}
)

// User will hold all the user information
type User struct {
	Nick     string
	Session  ssh.Session
	Term     *terminal.Terminal
	Color    string
	Addr     string
	Win      ssh.Window
	JoinTime time.Time
}

// New will instantiate and return a newly connected user
func New(s ssh.Session) *User {
	// Give terminal
	term := terminal.NewTerminal(s, fmt.Sprintf("%s $ ", s.User()))
	_ = term.SetSize(10000, 10000)
	pty, _, _ := s.Pty()
	w := pty.Window

	// Read remote IP
	// If error close session
	remoteIP, _, err := net.SplitHostPort(s.RemoteAddr().String())
	if err != nil {
		e := fmt.Sprintf("There has been an error connecting: %+v\n", err.Error())
		term.Write([]byte(e))
		s.Close()
		return nil
	}

	// TODO: maybe implement some kind of identification?
	// For Example client ssh pubkey?
	// Assign it to User ID then?

	return &User{
		Nick:     s.User(),
		Session:  s,
		Term:     term,
		Color:    "",
		Addr:     remoteIP,
		Win:      w,
		JoinTime: time.Now(),
	}
}

// ChangeColor will handle the change of terminal prompt color
func (u *User) ChangeColor(color string) error {
	return nil
}

// Print will display a text at users prompt
func (u *User) Print(sender, message string) {
	logger.Debugf("User Print function has received message '%s' from sender <%s>", sender, message)
	var msg string

	// Render Markdown
	if sender != "" { // someone is a sender
		msg = strings.TrimSpace(utils.MDRender(message, len(sender), u.Win.Width))
		msg = sender + msg + "\a"
	} else { // there is no sender (system messages and such)
		msg = strings.TrimSpace(utils.MDRender(message, 0, u.Win.Width))
	}

	// Write to users prompt
	u.Term.Write([]byte(msg + "\n"))
}

func (u *User) ChatLoop() {
	for {
		line, err := u.Term.ReadLine()
		if err == io.EOF {
			u.Session.Close()
			return
		}
		if err != nil {
			logger.Warnf("The chat loop of user %s had an error %+v", u.Nick, err)
			continue
		}
		line = strings.TrimSpace(line) // Strip Whitespaces at end

		u.Term.Write([]byte(line))

		logger.Debugf("Line was: %+v", line)

		// TODO Process message here
	}
}
