package server

import (
	"github.com/docker/docker/client"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var docker *client.Client

func NewServer() http.Handler {

	r := mux.NewRouter()
	r.HandleFunc("/project/{projectId}", HandleProject)
	r.HandleFunc("/project/{projectId}/builds", HandleBuilds)
	r.HandleFunc("/project/{projectId}/build/{buildVersion}", HandleBuild)
	r.HandleFunc("/project/{projectId}/build/{buildVersion}/tag/{tagType}", HandleTag)
	r.HandleFunc("/keys", HandleKeys)

	var err error
	docker, err = NewDockerClient("1.37")
	if err != nil {
		log.Fatalln(err)
	}

	return r
}
