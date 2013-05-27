package service

import (
	"bloom"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var filter bloom.Filter = bloom.New(100, 0.01)

func add(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[len("/add/"):]

	filter.Add([]byte(key))

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "/contains/%s", key)
}

func contains(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[len("/contains/"):]

	if filter.Contains([]byte(key)) {
		w.WriteHeader(http.StatusOK)

		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func createFilter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "createFilter%+v\n", mux.Vars(r))
	fmt.Fprintf(w, "name:%+v, capacity:%+v, errorRate:%+v\n", r.FormValue("name"), r.FormValue("capacity"), r.FormValue("errorRate"))
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

	r.HandleFunc("/v1/filters", createFilter).Methods("POST")
	r.HandleFunc("/v1/filters/{filterName}", retrieveFilter).Methods("GET")
	r.HandleFunc("/v1/filters/{filterName}", deleteFilter).Methods("DELETE")
	r.HandleFunc("/v1/filters/{filterName}/entries", createEntry).Methods("POST")
	r.HandleFunc("/v1/filters/{filterName}/entries/{entryName}", retrieveEntry).Methods("GET")

	http.Handle("/", r)

	log.Print("Server starting on ", address)
	http.ListenAndServe(address, nil)
}
