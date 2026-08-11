package main

import (
	"encoding/json"
	"net/http"
)

type Mahasiswa struct {
	Nama  string `json:"nama"`
	Nim   string `json:"nim"`
	Email string `json:"email"`
}

func getMahasiswa(w http.ResponseWriter, r *http.Request) {
	mahasiswa := Mahasiswa{
		Nama:  "Ilham",
		Nim:   "240411100009",
		Email: "ilhamgacor@gmail.com",
	}

	result, err := json.MarshalIndent(mahasiswa, "", " ")
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", getMahasiswa)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
