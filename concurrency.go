package concurrency

import (
	"fmt"
	"sync"
	"onefootball/business"
	"onefootball/model"
)

type Job struct {  
    id       int
}
type Result struct {  
    job         Job
	teamId int
	teamName string
}

const jobCount = 50
const resultCount = 10
var jobs = make(chan Job, jobCount)  
var results = make(chan Result, resultCount)


func Worker(mutex *sync.Mutex, wg *sync.WaitGroup, teamDetails *[]model.Output) {  
    for job := range jobs {
		team, isValid := business.GetTeam(job.id)
		if isValid {
            mutex.Lock() 
            *teamDetails = append(*teamDetails, team)
            mutex.Unlock()  
            
        	output := Result{job, team.Data.Team.Id, team.Data.Team.Name}
			results <- output
		}
    }
    wg.Done()
}
func CreateWorkerPool(teamDetails *[]model.Output) {  
    var wg sync.WaitGroup
    var m sync.Mutex
    for i := 0; i < jobCount; i++ {
        wg.Add(1)
        go Worker(&m, &wg, teamDetails)
    }
    wg.Wait()
    close(results)
}
func Allocate(noOfJobs int) {  
    for i := 0; i < noOfJobs; i++ {
        job := Job{i}
        jobs <- job
    }
    close(jobs)
}
func GetResult(done chan bool) {  
    i := 1
    fmt.Println("Founded teams:\n")
    for result := range results {
        fmt.Printf("%v: team name: %v ,team id: %v\n", i, result.teamName, result.teamId)
        i++
	}
    done <- true
}