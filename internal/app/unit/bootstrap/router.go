package bootstrap

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetRouter(handlers *HandlerContainer) (*mux.Router, error) {
	r := mux.NewRouter()
	r.HandleFunc("/speedtest", handlers.handler.Speedtest).Methods(http.MethodGet)
	return r, nil
}
