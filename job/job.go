package job

type JobBody struct {
    Command string  `json: Command line`
    MaintanerAddress string `json: Maintainer address`
}

type Job struct {
    Name string    `json: Jobname`
    Body JobBody   `json: Body`
}

type JobResult struct {
    JobItem Job
    StdOut string  `json: Stdout`
    StdErr string  `json: Stderr`
    Success bool   `json: Success`
}
