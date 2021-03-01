package util

import (
	"encoding/json"
	"io"
	"net/http"
)

func DecodeJSON(r io.Reader, data interface{}) ErrorResponse {
	err := json.NewDecoder(r).Decode(&data)

	if err != nil {
		return ErrorResponse{Code: "cant-parse-data", Status: 400}
	}
	return ErrorResponse{}
}

func HandleError(w http.ResponseWriter, err ErrorResponse) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(err.Status)
	json.NewEncoder(w).Encode(err)
}

func PrepareResponse(w http.ResponseWriter, data interface{}, err ErrorResponse) {
	if len(err.Code) > 0 {
		HandleError(w, err)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}
