package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	machinery "github.com/RichardKnop/machinery/v1"
	machineryConfig "github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
)

// Pointer utama untuk machinery, akan dipakai ulang di berbagai tempat
var server *machinery.Server

// Fungsi utama
func main() {

	// Inisialisasi machinery
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

	// Inisialisasi http service
	httpServices()

}

func httpServices() {
	http.HandleFunc("/submit", indexHandler)

	port := os.Getenv("PORT")
	if len(port) < 1 {
		port = "8080"
	}
	log.Println("irgsh-go chief now live on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}

// HTTP handler
func indexHandler(w http.ResponseWriter, r *http.Request) {

	ID := "123" // ID task
	// Buat tugas
	buildSignature := tasks.Signature{
		Name: "build", // Nama worker
		UUID: ID,
		Args: []tasks.Arg{
			{
				Type:  "string",
				Value: "foobar",
			},
		},
	}

	repoSignature := tasks.Signature{
		Name: "repo", // Nama worker
		UUID: ID,
	}

	// Masukkan tugas ke dalam rantai tugas
	chain, _ := tasks.NewChain(&buildSignature, &repoSignature)

	// Kirim tugas ke worker
	_, err := server.SendChain(chain)
	if err != nil {
		fmt.Println("Could not send chain : " + err.Error())
	}
	resp := "OK"
	fmt.Fprintf(w, resp)
}
