package main

import (
	"sync"

	"github.com/zacharygilliom/JobFinder/internal/authtoken"
	"github.com/zacharygilliom/JobFinder/internal/email"
)

func main() {
	serv, user := authtoken.ConnectClient()
	ml := email.GetMessageIDs(serv, user)
	ems := ml.GetMessages(serv, user, "From")
	//c := colly.NewCollector()
	/*
		for i := 0; i < len(ems.List); i++ {
			var wg sync.WaitGroup
			for j := 0; j < len(ems.List[i].Jobs); j++ {
				wg.Add(1)
				//job := ems.List[i].Jobs[j]
				go runParser(ems.List[i].Jobs[j], c, &wg)
				// go runParser(job, c)
			}
			wg.Wait()

		}
	*/
	for _, i := range ems.List {
		var wg sync.WaitGroup
		for _, j := range i.Jobs {
			wg.Add(1)
			job := j
			go runParser(job, &wg)

		}
		wg.Wait()
	}
	//fmt.Println("Reading Data from Websites into Channel")

	//ems.GetJobInfo()
	//fmt.Println("Finished Reading Data into Channel --- Channel Closed")
}

func runParser(j email.Job, wg *sync.WaitGroup) {
	defer wg.Done()
	j.ParseSite()
}
