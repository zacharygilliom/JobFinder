package main

import (
	"github.com/zacharygilliom/JobFinder/internal/authtoken"
	"github.com/zacharygilliom/JobFinder/internal/email"
)

func main() {
	serv, user := authtoken.ConnectClient()
	ml := emails.GetMessageIDs(serv, user)
	ml.GetMessages(serv, user, "From")
}
