package utils

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Errors  any    `json:"errors,omitempty"`
	w       http.ResponseWriter
}

type Response struct {
	Code    int
	Data    any
	Type    string
	Message string
}

func New(w http.ResponseWriter) *response {
	w.Header().Set("Content-Type", "application/json")

	return &response{
		w: w,
	}
}

func (r *response) WithInternalServerError() {
	r.Code = http.StatusInternalServerError
	r.Message = "internal server error"

	r.w.WriteHeader(r.Code)

	err := json.NewEncoder(r.w).Encode(r)
	if err != nil {
		return
	}
}

func (r *response) WithData(ctr *Response) {
	if ctr.Data != nil {
		r.Data = ctr.Data
	}
	err := json.NewEncoder(r.w).Encode(r)

	if err != nil {
		return
	}
}
