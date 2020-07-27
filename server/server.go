package main

import (
	"encoding/json"
	"net/http"
	"time"

	"log"
)

type timeHandler struct{}

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

func main() {
	h := &timeHandler{}
	log.Fatal(http.ListenAndServe(":8080", h))
}
