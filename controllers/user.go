package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"nshare-api/models"
	"nshare-api/responses"
)

func POSTInitSession(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		json.NewDecoder(r.Body).Decode(&user)

		user.Passhash = user.CreatePassHash().Passhash

		rows, err := db.Query("SELECT * FROM nsh_users WHERE `passhash`=?", user.Passhash)

		if err != nil {
			panic(err)
		}

		var usr models.User
		var usrsDump []models.User

		defer rows.Close()
		for rows.Next() {
			rows.Scan(&usr.ID, &usr.Login, &usr.Passhash, &usr.Permissions, &usr.SA)
			usrsDump = append(usrsDump, usr)
		}

		if len(usrsDump) == 0 {
			responses.SendUnSucResp(w, 401, "Bad password or login")
		} else {
			var session models.Session
			session.SKey = usrsDump[0].InitSession(db)
			session.UserID = usrsDump[0].ID

			responses.SendSucResp(w, session)
		}
	}
}
