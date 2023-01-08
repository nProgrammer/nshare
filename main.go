package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	sslmode, err := strconv.ParseBool(os.Getenv("SSLMODE"))

	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})

	if sslmode {
		log.Fatal(http.ListenAndServeTLS(":8080", "certificates/origin.pem", "certificates/private.key", r))
	} else {
		log.Fatal(http.ListenAndServe(":"+os.Getenv("WORKINGPORT"), r))
	}
}
