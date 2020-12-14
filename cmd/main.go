package main

import (
	"fmt"
	"time"

	"github.com/zacharygilliom/JobFinder/internal/authtoken"
	"github.com/zacharygilliom/JobFinder/internal/email"
)

func main() {
	serv, user := authtoken.ConnectClient()
	ml := email.GetMessageIDs(serv, user)
	ems := ml.GetMessages(serv, user, "From")
	c := make(chan string)
	fmt.Println("Reading Data from Websites into Channel")
	t0 := time.Now()
	go ems.GetJobInfo(c)
	//close(c)
	for _, val := range c {

	}
	t1 := time.Now()
	fmt.Printf("Time taken to parse websites: %v\n", t1.Sub(t0))
	fmt.Println("Finished Reading Data into Channel --- Channel Closed")
}
