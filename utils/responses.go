package utils

import (
    "encoding/json"
    "net/http"
)

type ErrorResponse struct {
    Message string `json:"message"`
    Code    int    `json:"code"`
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
    response, _ := json.Marshal(ErrorResponse{Message: message, Code: code})
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}
