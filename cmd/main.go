package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/zacharygilliom/JobFinder/internal/authtoken"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

func main() {
	authtoken.Verify()
	getMessages()
}

func getMessages() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	gmailService, err := gmail.NewService(ctx, option.WithCredentialsFile("../credentials.json"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gmailService)

	messagesCall := gmailService.Users.Messages.List("me")
	messagesResp, err := messagesCall.Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messagesResp)

}
