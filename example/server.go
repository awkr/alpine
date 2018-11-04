package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("recvd req")
		fmt.Fprintf(w, "hi, it is %+v.", time.Now())
	})
	log.Println("server started")
	http.ListenAndServe(":7070", mux)
}
