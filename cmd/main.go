package main

import (
	"fmt"

	"github.com/zacharygilliom/JobFinder/internal/authtoken"
	"github.com/zacharygilliom/JobFinder/internal/email"
)

func main() {
	serv, user := authtoken.ConnectClient()
	ml := email.GetMessageIDs(serv, user)
	ems := ml.GetMessages(serv, user, "From")
	c := make(chan string)

	fmt.Println("Reading Data from Websites into Channel")
	go ems.GetJobInfo(c)

	fmt.Println("Finished Reading Data into Channel --- Channel Closed")
}
