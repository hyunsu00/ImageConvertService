package main

import (
	"imgconv/api"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/convert", api.ConvertHandler)

	log.Println("서버 시작: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
