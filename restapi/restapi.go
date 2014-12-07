package restapi

import (
    "encoding/json"
    "log"
    "net/http"
    "regexp"
    "strings"
    "github.com/bgabor666/go-jobserver/job"
    "github.com/bgabor666/go-jobserver/jobserver"
)


type RestAPI struct {
    jobServer *jobserver.JobServer
}

func NewRestAPI(js *jobserver.JobServer, port string) *RestAPI {
    ra := &RestAPI{jobServer: js}
    err := http.ListenAndServe(port, ra)
    if err != nil {
        panic(err)
    }
    return ra
}


func (ra *RestAPI) ServeHTTP(writer http.ResponseWriter, request* http.Request) {
    var validJobURL = regexp.MustCompile(`^\/\:[a-zA-Z0-9]+$`)
    var validLastRunURL = regexp.MustCompile(`^\/\:[a-zA-Z0-9]+/:lastrun$`)
    if request.Method == "GET" {
        if request.URL.Path == "/" {
            jobListAsJSON, err := json.Marshal(ra.jobServer.JobList)
            if err != nil {
                http.Error(writer, err.Error(), http.StatusInternalServerError)
                return
            }
            writer.Header().Set("Content-Type", "application/json")
            writer.Write(jobListAsJSON)
            return
        } else if validJobURL.MatchString(request.URL.Path) {
            jobName := request.URL.Path[2:]
            if body, ok := ra.jobServer.JobList[jobName]; ok {
                job := job.Job{Name: jobName, Body: body}
                jobAsJSON, err := json.Marshal(job)
                if err != nil {
                    http.Error(writer, err.Error(), http.StatusInternalServerError)
                    return
                }
                writer.Header().Set("Content-Type", "application/json")
                writer.Write(jobAsJSON)
                return
	    }
        } else if validLastRunURL.MatchString(request.URL.Path) {
            jobName := strings.Split(request.URL.Path[2:], "/")[0]
            if ra.jobServer.HasLastResult(jobName) {
                resultAsJSON, err := json.Marshal(ra.jobServer.GetLastResult(jobName))
		log.Println(ra.jobServer.GetLastResult(jobName))
                if err != nil {
                    http.Error(writer, err.Error(), http.StatusInternalServerError)
                    return
                }
                writer.Header().Set("Content-Type", "application/json")
                writer.Write(resultAsJSON)
                return
	    }
	}
    } else if request.Method == "POST" {
        if validJobURL.MatchString(request.URL.Path) {
            jobName := request.URL.Path[2:]
            if _, ok := ra.jobServer.JobList[jobName]; ok {
                go ra.jobServer.StartJob(jobName)
                return
	    }
        }
    }
    writer.WriteHeader(http.StatusNotFound)
}