package main

import (
	"fmt"
	"log"

	machinery "github.com/RichardKnop/machinery/v1"
	machineryConfig "github.com/RichardKnop/machinery/v1/config"
)

var server *machinery.Server

func main() {
	svr, err := machinery.NewServer(
		&machineryConfig.Config{
			Broker:        "redis://localhost:6379",
			ResultBackend: "redis://localhost:6379",
			DefaultQueue:  "irgsh",
		},
	)
	if err != nil {
		fmt.Println("Could not create server : " + err.Error())
	}
	server = svr

	server.RegisterTask("repo", Repo)
	worker := server.NewWorker("repo", 1)
	err = worker.Launch()
	if err != nil {
		fmt.Println("Could not launch worker : " + err.Error())
	}
}

func Repo(payload string) (next string, err error) {
	log.Println(payload)
	return payload, nil
}
