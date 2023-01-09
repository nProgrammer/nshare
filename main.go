package main

import (
	"log"
	"net/http"
	"nshare-api/controllers"
	"nshare-api/models"
	"nshare-api/responses"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	sslmode, err := strconv.ParseBool(os.Getenv("SSLMODE"))

	var dbConnector models.DBConnector
	dbConnector = dbConnector.PrepareVars()
	dbConnector.Connection = dbConnector.ConnectDB().Connection

	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		responses.SendSucResp(w, "Works")
	}).Methods("GET")

	r.HandleFunc("/users/initSession", controllers.POSTInitSession(dbConnector.Connection)).Methods("POST")

	if sslmode {
		log.Fatal(http.ListenAndServeTLS(":8080", "certificates/origin.pem", "certificates/private.key", r))
	} else {
		log.Fatal(http.ListenAndServe(":"+os.Getenv("WORKINGPORT"), r))
	}
}
