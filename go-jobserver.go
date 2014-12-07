package main

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "github.com/bgabor666/go-jobserver/jobserver"
    "github.com/bgabor666/go-jobserver/restapi"
)

func main() {
    server := jobserver.NewJobServer()

    data, err := ioutil.ReadFile("joblist.yaml")
    if err != nil {
        panic("Cannot open joblist.yaml: " + err.Error())
    }
    err = yaml.Unmarshal(data, &server.JobList)
    if err != nil {
        panic("Cannot load joblist.yaml: " + err.Error())
    }

   restapi.NewRestAPI(server, ":8060")
}
