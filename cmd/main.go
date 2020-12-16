package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/zacharygilliom/JobFinder/internal/authtoken"
	"github.com/zacharygilliom/JobFinder/internal/email"
)

func main() {
	serv, user := authtoken.ConnectClient()
	ml := email.GetMessageIDs(serv, user)
	ems := ml.GetMessages(serv, user, "From")
	c := colly.NewCollector()
	fmt.Println("Reading Data from Websites into Channel")

	ems.GetJobInfo(c)
	fmt.Println("Finished Reading Data into Channel --- Channel Closed")
}
