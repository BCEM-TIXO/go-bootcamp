package main

import (
	"log"
	"net/http"
	"os"

	// "ex02/server/database"
	sw "ex02/server/handlers"
)

func main() {

	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	app := sw.App{}
	app.Initialize(logger)

	log.Fatal(http.ListenAndServe(":8080", app.Router))
	log.Println("Server started on :8080")
}
