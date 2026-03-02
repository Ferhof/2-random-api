package main

import (
	"math/rand"
	"net/http"
	"strconv"
)

func randomHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	w.Write([]byte(strconv.Itoa(rand.Intn(6) + 1)))
}

func main() {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	mux.HandleFunc("/random", randomHandler)
	server.ListenAndServe()
}
