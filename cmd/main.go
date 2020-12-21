package main

import (
	"github.com/zacharygilliom/JobFinder/internal/authtoken"
	"github.com/zacharygilliom/JobFinder/internal/email"
)

func main() {
	serv, user := authtoken.ConnectClient()
	ml := email.GetMessageIDs(serv, user)
	ems := ml.GetMessages(serv, user, "From")
	//var parsedEms email.Emails
	for _, i := range ems.List {
		//var parsedEm email.Email
		//var wg sync.WaitGroup
		for _, j := range i.Jobs {
			//wg.Add(1)
			job := j
			job.ParseSite()
			//parsedEm.Jobs = append(parsedEm.Jobs, job)
		}
		//parsedEms.List = append(parsedEms.List, parsedEm)
	}
	//fmt.Println("Reading Data from Websites into Channel")

	//parsedEms.GetJobInfo()
	ems.GetJobInfo()
	//fmt.Println(parsedEms.List)
	//fmt.Println("Finished Reading Data into Channel --- Channel Closed")
}

func runParser(j email.Job) {
	j.ParseSite()
}
