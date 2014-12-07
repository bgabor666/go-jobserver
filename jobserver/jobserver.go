package jobserver

import (
    "log"
    "sync"
    "github.com/bgabor666/go-jobserver/database"
    "github.com/bgabor666/go-jobserver/job"
    "github.com/bgabor666/go-jobserver/worker"
)

type JobServer struct {
    JobList map[string]job.JobBody
    DataStorage *database.DataBackend
}


func (js *JobServer) StartJob(jobName string) {
    var result job.JobResult
    var wg sync.WaitGroup
    job := job.Job{Name: jobName, Body: js.JobList[jobName]}
    w := new(worker.Worker)

    wg.Add(1)
    go w.StartJob(job, &result, &wg)
    wg.Wait()

    go js.DataStorage.StoreResult(jobName, result)

    log.Println("JOBNAME: ", result.JobItem.Name)
    if result.Success {
        log.Println("SUCCESS: ", "success")
    } else {
        log.Println("SUCCESS: ", "failure")
    }
    log.Println("STDOUT:\n", result.StdOut)
    log.Println("STDERR:\n", result.StdErr)
}

func (js *JobServer) HasLastResult(jobName string) bool {
    if _, ok := js.DataStorage.LastRun[jobName]; ok {
        return true
    }
    return false
}

func (js *JobServer) GetLastResult(jobName string) job.JobResult {
    return js.DataStorage.LastRun[jobName]
}

func NewJobServer() *JobServer {
    return &JobServer{JobList: make(map[string]job.JobBody), DataStorage: database.NewDataBackend()}
}
