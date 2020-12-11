package emails

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"google.golang.org/api/gmail/v1"
)

// Email ...
type Email struct {
	From string
	URL  string
}

// Message ...
type Message struct {
	ID       string `json:"id"`
	ThreadID string `json:"threadId"`
}

// MessageList ...
type MessageList struct {
	Messages []Message `json:"messages"`
}

// GetMessageIDs ...
func GetMessageIDs(serv *gmail.Service, user string) MessageList {
	messages, err := serv.Users.Messages.List(user).Do()
	if err != nil {
		log.Fatal(err)
	}
	messageDetails, err := messages.MarshalJSON()
	var ml MessageList
	json.Unmarshal(messageDetails, &ml)

	return ml
}

// GetMessages ...
func (ml *MessageList) GetMessages(serv *gmail.Service, user string, header string) {
	for _, mess := range ml.Messages {
		messages, err := serv.Users.Messages.Get(user, mess.ID).Do()
		if err != nil {
			log.Fatal(err)
		}
		var em Email
		encodedMessage := messages.Payload.Body.Data
		encodedHeaders := messages.Payload.Headers
		encodedData := messages.Payload.Parts
		fmt.Println(encodedData)
		for _, head := range encodedHeaders {
			fmt.Println(head)
			if head.Name == header {
				em.From = head.Value
			}
		}
		if strings.Contains(em.From, "alert@indeed.com") {
			fmt.Println(em.From)
			fmt.Println(encodedMessage)
			byteMessage, err := base64.URLEncoding.DecodeString(encodedMessage)
			for _, part := range encodedData {
				stringPart, err := base64.URLEncoding.DecodeString(part.Body.Data)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(string(stringPart))
			}
			if err != nil {
				fmt.Printf("Error decoding string: %s ", err.Error())
				return
			}

			fmt.Println("*********************************************************  New Message **********************************************")
			fmt.Println(byteMessage)
			fmt.Println("*********************************************************  End Of Message *******************************************")
		}

	}

}
