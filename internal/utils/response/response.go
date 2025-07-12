package response

import (
	"encoding/json"
	"net/http"
)

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

type Response struct {
	Status  string                 `json: "status"`
	Message string                 `json:"message,omitempty"`
	Error   string                 `json:"error,omitempty"`
	Errors  map[string]string      `json:"errors,omitempty"` // field-level validation errors
	Data    map[string]interface{} `json:"data,omitempty"`   // optional for success payload
}

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)

}

func GeneralError(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

// ValidationError returns a response specifically for field validation issues
func ValidationError(fields map[string]string) Response {
	return Response{
		Status:  StatusError,
		Message: "Validation failed",
		Errors:  fields,
	}
}
