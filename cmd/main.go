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
	var parsedEms email.Emails
	for _, i := range ems.List {
		var parsedEm email.Email
		for _, j := range i.Jobs {
			var job email.Job
			jo := j
			job = email.ParseSite(jo)
			//fmt.Printf("Job title: %v\nDoes it Include a keyword: %v\n", job.Title, job.Valid)
			//fmt.Printf("%+v\n", job)
			parsedEm.Jobs = append(parsedEm.Jobs, job)
		}
		parsedEms.List = append(parsedEms.List, parsedEm)
	}
	//fmt.Println("Reading Data from Websites into Channel")

	//parsedEms.GetJobInfo()
	//ems.GetJobInfo()
	fmt.Println("**************************************************\n*******************\n**************************")
	parsedEms.GetJobInfo()
	// fmt.Println(parsedEms.List)
	//fmt.Println("Finished Reading Data into Channel --- Channel Closed")
}
