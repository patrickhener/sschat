package rooms

import (
	"sync"

	"github.com/patrickhener/sschat/chatuser"
	"github.com/patrickhener/sschat/logger"
)

var (
	// Init default room 'general'
	GeneralRoom = &Room{
		Name:       "#general",
		Users:      make([]*chatuser.User, 0, 10),
		UsersMutex: sync.Mutex{},
	}

	// Map of rooms
	Rooms = map[string]*Room{
		GeneralRoom.Name: GeneralRoom,
	}
)

// Room represents a chatroom
type Room struct {
	Name       string
	Users      []*chatuser.User
	UsersMutex sync.Mutex
}

func (r *Room) Broadcast(sender, message string) {
	logger.Debugf("Broadcast received message '%s' from sender <%s>", message, sender)
	// No empty messages
	if message == "" {
		return
	}

	// Print to all users prompts in room
	r.UsersMutex.Lock()
	for i := range r.Users {
		r.Users[i].Print(sender, message)
	}
	r.UsersMutex.Unlock()
}
