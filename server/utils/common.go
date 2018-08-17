package utils

import (
	"encoding/json"
	"net/http"
)

type commResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Reply can be used for replying response
// Rename this to "reply" in the future because should call ResponseReply instead
func Reply(w http.ResponseWriter, r *http.Request, status int, v interface{}) {
	data, _ := json.Marshal(v)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}

// CommReply can be used for replying some common data
func CommReply(w http.ResponseWriter, r *http.Request, status int, message string) {
	resp := commResp{
		Code:    http.StatusText(status),
		Message: message,
	}
	Reply(w, r, status, resp)
}
