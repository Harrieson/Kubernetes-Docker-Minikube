package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	started := time.Now()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		duration := time.Now().Sub(started)

		if duration.Seconds() > 1000 {
			log.Println("TimeOut Triggered")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`Timed Out Now!`))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`Hello Gopher`))
		}
		fmt.Fprintf(w, "Hello Gophers")
	})

	http.HandleFunc("/aligator", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Gator")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
