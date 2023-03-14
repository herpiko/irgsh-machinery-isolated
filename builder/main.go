package main

import (
	"fmt"
	"log"
	"time"

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

	server.RegisterTask("build", Build)
	worker := server.NewWorker("builder", 1)
	err = worker.Launch()
	if err != nil {
		fmt.Println("Could not launch worker : " + err.Error())
	}
}

func Build(payload string) (next string, err error) {
	log.Println(payload)
	time.Sleep(5 * time.Second)
	return payload, nil
}
