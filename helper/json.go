package helper

import (
	"encoding/json"
	"net/http"
)

func RequestFromStruct(r *http.Request, result any) {
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(result)
	if err != nil {
		panic(err)
	}
}

func WriteToResponse(w http.ResponseWriter, response any) {

	w.Header().Add("content-type", "application/json")
	encode := json.NewEncoder(w)
	err := encode.Encode(response)
	if err != nil {
		panic(err)
	}
}
