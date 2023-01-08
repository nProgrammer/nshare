package models

import (
	"database/sql"
	"nshare-api/utils"
	"strings"
)

type User struct {
	ID          int    `json:"id"`
	Login       string `json:"login"`
	Passhash    string `json:"passhash"`
	Permissions string `json:"permissions"`
}

func CheckPermissions(rwd string, do string) bool {
	permissions := strings.Split(rwd, "")
	i := 0
	for i < 3 {
		if permissions[i] == do {
			return true
		}
		i++
	}
	return false
}

func (usr User) InitSession(db *sql.DB) string {
	skey := utils.RandStringRunes(11)
	skey = utils.Sha256(skey)
	db.Exec("INSERT INTO `nsh_sessions` (userid, skey) VALUE (?,?)", usr.ID, skey)

	return skey
}
