package service

import (
    "bloom"
    "fmt"
    "net/http"
    "log"
)

var filter bloom.Filter = bloom.New(100,0.01)

func add(w http.ResponseWriter, r *http.Request) {
    key := r.URL.Path[len("/add/"):]

    filter.Add([]byte(key));

    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "/contains/%s", key)
}

func contains(w http.ResponseWriter, r *http.Request) {
    key := r.URL.Path[len("/contains/"):]

    if (filter.Contains([]byte(key))) {
        w.WriteHeader(http.StatusOK)

        return
    }

    w.WriteHeader(http.StatusNotFound)
}

func ListenAndServe(address string) {
    http.HandleFunc("/add/", add)
    http.HandleFunc("/contains/", contains)
    log.Print("Server starting on ", address)
    http.ListenAndServe(address, nil)
}
