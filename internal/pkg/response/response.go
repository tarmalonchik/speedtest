package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// ErrorResponse - error skeleton
type ErrorResponse struct {
	Error      string `json:"error"`
	StatusCode int    `json:"statusCode"`
}

// SuccessfulResponse - default successful message
type SuccessfulResponse struct {
	Result     string `json:"result"`
	StatusCode int    `json:"statusCode"`
}

func JSON200Tinkoff(w http.ResponseWriter) {
	_, _ = w.Write([]byte("OK"))
}

func JSON200(w http.ResponseWriter) {
	data := "OK"
	RenderJSON(w, http.StatusOK, data)
}

// JSON500 - shortcut for JSON500txt
func JSON500(w http.ResponseWriter) {
	JSON500txt(w, "internal server error")
}

// JSON500txt - responds with 500 error
func JSON500txt(w http.ResponseWriter, errText string) {
	data := ErrorResponse{
		Error:      errText,
		StatusCode: http.StatusInternalServerError,
	}
	RenderJSON(w, http.StatusInternalServerError, data)
}

// JSON400 - shortcut for JSON400txt
func JSON400(w http.ResponseWriter) {
	JSON400txt(w, "bad request")
}

// JSON404 - shortcut for JSON404txt
func JSON404(w http.ResponseWriter) {
	JSON404txt(w, "not found")
}

// JSON400txt - 400 response with error text
func JSON400txt(w http.ResponseWriter, errText string) {
	data := ErrorResponse{
		Error:      errText,
		StatusCode: http.StatusBadRequest,
	}
	RenderJSON(w, http.StatusBadRequest, data)
}

// JSON404txt - 404 response with error text
func JSON404txt(w http.ResponseWriter, errText string) {
	data := ErrorResponse{
		Error:      errText,
		StatusCode: http.StatusNotFound,
	}
	RenderJSON(w, http.StatusNotFound, data)
}

// RenderJSON generic json response
func RenderJSON(w http.ResponseWriter, status int, data interface{}) {
	resp, err := json.Marshal(data)
	if err != nil {
		log.WithField("data", data).Error("Could not encode response")
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(resp)
	if err != nil {
		log.WithError(err).Error("RenderJSON w.Write error")
	}
}

func RenderString(w http.ResponseWriter, status int, data interface{}) {
	resp := []byte(fmt.Sprint(data))
	w.Header().Set("Content-type", "text/plain")
	w.WriteHeader(status)
	_, err := w.Write(resp)
	if err != nil {
		log.WithError(err).Error("RenderJSON w.Write error")
	}
}
