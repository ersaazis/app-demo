package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	phpURL := os.Getenv("PHP_APP_URL")
	if phpURL == "" {
		phpURL = "http://php-app:8000"
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
	})

	http.HandleFunc("/some-error", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal Server Error")
	})

	http.HandleFunc("/some-load", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get(phpURL + "/go")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		w.WriteHeader(resp.StatusCode)
		w.Write(body)
	})

	http.HandleFunc("/some-call-error", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get(phpURL + "/query")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		w.WriteHeader(resp.StatusCode)
		w.Write(body)
	})

	fmt.Println("Go app starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
