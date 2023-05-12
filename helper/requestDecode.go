package helper

import (
	"encoding/json"
	"net/http"
)

func RequestDecode(r *http.Request, data interface{}) {
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(data)
	ErrorHandle(err)
}
