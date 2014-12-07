package main

import (
    "github.com/bgabor666/go-jobserver/job"
    "github.com/bgabor666/go-jobserver/jobserver"
    "github.com/bgabor666/go-jobserver/restapi"
)

func main() {
    job1 := job.Job{Name: "configure", Body: job.JobBody{Command: "./configure", MaintanerAddress: "maintaner@address.com"}}
    job2 := job.Job{Name: "build", Body: job.JobBody{Command: "make", MaintanerAddress: "maintaner@address.com"}}
    job3 := job.Job{Name: "ls", Body: job.JobBody{Command: "ls -l .. && ls -l", MaintanerAddress: "maintaner@address.com"}}
    job4 := job.Job{Name: "sleep5", Body: job.JobBody{Command: "sleep 5", MaintanerAddress: "maintaner@address.com"}}

    server := jobserver.NewJobServer()
    server.AddJob(job1)
    server.AddJob(job2)
    server.AddJob(job3)
    server.AddJob(job4)

    restapi.NewRestAPI(server, ":8060")
}
