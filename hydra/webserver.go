package main

import (
	"fmt"
	"net/http"

	"github.com/gy-kim/mastering-go-programming/hydra/hlogger"
)

func main() {
	logger := hlogger.GetInstance()
	logger.Println("Starting Hydra web service")
	http.HandleFunc("/", sroot)
	http.ListenAndServe(":8000", nil)
}

func sroot(w http.ResponseWriter, r *http.Request) {
	logger := hlogger.GetInstance()
	fmt.Fprintf(w, "Welcome to the Hydra software system")

	logger.Println("Received an http Get request on root url")
}
