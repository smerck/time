package main

import (
	"encoding/json"
	"net/http"
	"time"

	"log"

	"github.com/gorilla/mux"
)

type timeHandler struct{}
type healthHandler struct{}

type timeResponse struct {
	Time string `json:"time"`
}

func (h *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC3339)
	resp := &timeResponse{
		Time: t,
	}

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Println("failed to encode response: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *healthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/healthcheck").Handler(&healthHandler{})
	r.PathPrefix("/").Handler(&timeHandler{})
	log.Fatal(http.ListenAndServe(":9001", r))
}
