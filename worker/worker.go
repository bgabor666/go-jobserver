package worker

import (
    "bytes"
    "os/exec"
    "sync"
    "github.com/bgabor666/go-jobserver/job"
)

type Worker struct {
    shellCommand exec.Cmd
}

func (worker *Worker) StartJob(j job.Job, result *job.JobResult, wg *sync.WaitGroup) {
    defer wg.Done()
    var stdOutBuffer bytes.Buffer
    var stdErrBuffer bytes.Buffer
    result.JobItem = j

    worker.shellCommand = *exec.Command("bash", "-c", j.Body.Command)
    worker.shellCommand.Stdout = &stdOutBuffer
    worker.shellCommand.Stderr = &stdErrBuffer
    worker.shellCommand.Run()
    result.StdOut = stdOutBuffer.String()
    result.StdErr = stdErrBuffer.String()
    result.Success = worker.shellCommand.ProcessState.Success()
}