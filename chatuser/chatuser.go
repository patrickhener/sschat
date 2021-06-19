package chatuser

import "github.com/gliderlabs/ssh"

// User will hold all the user information
type User struct {
	Nick string
}

// New will instantiate and return a newly connected user
func New(s ssh.Session) *User {
	var u User
	u.Nick = s.User()
	return &u
}
