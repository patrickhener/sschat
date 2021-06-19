package main

import (
	"github.com/patrickhener/sschat/sschat"
)

func main() {
	port := 8000
	ip := "127.0.0.1"
	sshkey := "/home/patrick/.ssh/id_rsa"

	chat := sschat.Server{
		Port:   port,
		IP:     ip,
		SSHKey: sshkey,
	}

	chat.Run()
}
