package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/zacharygilliom/JobFinder/internal/authtoken"
	"google.golang.org/api/gmail/v1"
)

type message struct {
	ID       string `json:"id"`
	ThreadID string `json:"threadId"`
}
type messageList struct {
	Messages []message `json:"messages"`
}

func main() {
	serv, user := authtoken.ConnectClient()
	ml := getMessages(serv, user)
	fmt.Println(ml.Messages)
}

func getMessages(serv *gmail.Service, user string) messageList {
	messages, err := serv.Users.Messages.List(user).Do()
	if err != nil {
		log.Fatal(err)
	}
	messageDetails, err := messages.MarshalJSON()
	var ml messageList
	json.Unmarshal(messageDetails, &ml)

	return ml
}
