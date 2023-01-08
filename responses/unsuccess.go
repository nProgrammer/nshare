package responses

import (
	"encoding/json"
	"net/http"
)

func SendUnSucResp(w http.ResponseWriter, errNum int, x interface{}) {
	w.WriteHeader(errNum)
	json.NewEncoder(w).Encode(x)
}
