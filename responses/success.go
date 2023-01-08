package responses

import (
	"encoding/json"
	"net/http"
)

func SendSucResp(w http.ResponseWriter, x interface{}) {
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(x)
}
