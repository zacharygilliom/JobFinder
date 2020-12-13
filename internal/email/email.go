package email

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil" "log" "strings"

	"net/http"

	"google.golang.org/api/gmail/v1"
)

// Job ...
type Job struct {
	Title       string
	Location    string
	Description string
	URL         string
}

// Email ...
type Email struct {
	From string
	Jobs []Job
}

// Emails ...
type Emails struct {
	List []Email
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
func (ml *MessageList) GetMessages(serv *gmail.Service, user string, header string) Emails {
	var ems Emails
	for _, message := range ml.Messages {
		mess, err := serv.Users.Messages.Get(user, message.ID).Do()
		if err != nil {
			log.Fatal(err)
		}
		Headers := mess.Payload.Headers
		encodedPartsData := mess.Payload.Parts
		for _, head := range Headers {
			if head.Name == header {
				if strings.Contains(head.Value, "alert@indeed.com") {
					var em Email
					em.From = "Indeed"
					jobsPart, err := base64.URLEncoding.DecodeString(encodedPartsData[0].Body.Data)
					jobsPartString := string(jobsPart)
					if err != nil {
						log.Fatal(err)
					}
					em.GetJobsURL(jobsPartString)
					ems.List = append(ems.List, em)
				}
			}
		}
	}
	//fmt.Println(ems.List)
	return ems
}

// GetJobsURL ...
func (em *Email) GetJobsURL(emailBody string) {
	separatedStrings := strings.Split(emailBody, "\n")
	for _, val := range separatedStrings {
		trimmedURL := strings.TrimSuffix(val, "\r")
		var j Job
		if strings.Contains(trimmedURL, "https://www.indeed.com/rc/clk/") {
			j.URL = trimmedURL
			em.Jobs = append(em.Jobs, j)
		}

	}
}

// ParseSite ...
func (j *Job) ParseSite() string {
	resp, err := http.Get(j.URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}

// GetJobInfo ...
func (ems *Emails) GetJobInfo(c chan string) {
	for _, em := range ems.List {
		for _, j := range em.Jobs {
			c <- j.ParseSite()
		}
	}
	close(c)
}