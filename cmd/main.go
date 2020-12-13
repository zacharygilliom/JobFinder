package main

import (
	"github.com/zacharygilliom/JobFinder/internal/authtoken"
	"github.com/zacharygilliom/JobFinder/internal/email"
)

func main() {
	serv, user := authtoken.ConnectClient()
	ml := email.GetMessageIDs(serv, user)
	ems := ml.GetMessages(serv, user, "From")
	ems.GetJobInfo()
}
