package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/loosec9n/Project_One_Web_Server/routes"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile("ProjectOneWebServer.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	http.HandleFunc("/form", routes.FormHandler)
	http.HandleFunc("/hello", routes.HelloHandler)

	fmt.Println("Starting serer at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
