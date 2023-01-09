package main

import (
	"fmt"
	"log"
	"net/http"
	"nshare-api/controllers"
	"nshare-api/models"
	"nshare-api/responses"
	"nshare-api/utils"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	var dataDirPath string = "/var/nshare-data"
	isExisting := utils.CheckDataDir(dataDirPath)
	if isExisting {
		sslmode, err := strconv.ParseBool(os.Getenv("SSLMODE"))

		var dbConnector models.DBConnector
		dbConnector = dbConnector.PrepareVars()
		dbConnector.Connection = dbConnector.ConnectDB().Connection

		if err != nil {
			panic(err)
		}

		r := mux.NewRouter()

		r.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
			if utils.CheckDataDir(dataDirPath) {
				responses.SendSucResp(w, "Works")
			} else {
				responses.SendUnSucResp(w, 500, utils.FBErr(dataDirPath+" directory doesn't exists", "Internal Server Error"))
			}
		}).Methods("GET")

		r.HandleFunc("/users/initSession", controllers.POSTInitSession(dbConnector.Connection)).Methods("POST")

		if sslmode {
			log.Fatal(http.ListenAndServeTLS(":8080", "certificates/origin.pem", "certificates/private.key", r))
		} else {
			log.Fatal(http.ListenAndServe(":"+os.Getenv("WORKINGPORT"), r))
		}
	} else {
		fmt.Println("Error: /var/nshare-data directory doesn't exist")
	}
}
