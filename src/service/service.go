package service

import (
    "bloom"
    "fmt"
    "net/http"
    "log"
    "github.com/gorilla/mux"
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

func createFilter(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "createFilter%+v\n", mux.Vars(r))
}

func retrieveFilter(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "retrieveFilter %+v\n", mux.Vars(r))
}

func deleteFilter(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "deleteFilter %+v\n", mux.Vars(r))
}

func createEntry(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "createEntry %+v\n", mux.Vars(r))
}

func retrieveEntry(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "retrieveEntry %+v\n", mux.Vars(r))
}

func ListenAndServe(address string) {
    r := mux.NewRouter()
    s := r.PathPrefix("/v1/filters").Subrouter()

    s.HandleFunc("", createFilter).Methods("POST")
    s.HandleFunc("/{filterName}", retrieveFilter).Methods("GET")
    s.HandleFunc("/{filterName}", deleteFilter).Methods("DELETE")
    s.HandleFunc("/{filterName}/entries", createEntry).Methods("POST")
    s.HandleFunc("/{filterName}/entries/{entryName}", retrieveEntry).Methods("GET")

    http.Handle("/", r)
    log.Print("Server starting on ", address)
    http.ListenAndServe(address, nil)
}
