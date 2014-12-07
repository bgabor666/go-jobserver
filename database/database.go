package database

import (
    "github.com/bgabor666/go-jobserver/job"
)

type DataBackend struct {
    LastRun map[string]job.JobResult
}

func NewDataBackend() *DataBackend {
    return &DataBackend{LastRun: make(map[string]job.JobResult)}
}

func (db *DataBackend) StoreResult(jobName string, jr job.JobResult) {
    db.LastRun[jobName] = jr
}
