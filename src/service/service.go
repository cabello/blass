package service

import (
	"bloom"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

var filters map[string]bloom.Filter = make(map[string]bloom.Filter)

func createFilter(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	if _, ok := filters[name]; ok {
		w.WriteHeader(http.StatusConflict)

		return
	}

	parsedCapacity, _ := strconv.ParseInt(r.FormValue("capacity"), 10, 0)
	capacity := int(parsedCapacity)
	errorRate, _ := strconv.ParseFloat(r.FormValue("errorRate"), 64)

	filters[name] = bloom.New(capacity, errorRate)
	w.WriteHeader(http.StatusCreated)
}

func retrieveFilter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["filterName"]

	filter, ok := filters[name]
	if !ok {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "{ \"capacity\": %+v, \"errorRate\": %+v }\n", filter.Capacity, filter.ErrorRate)
}

func deleteFilter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["filterName"]

	if _, ok := filters[name]; !ok {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	delete(filters, name)
	w.WriteHeader(http.StatusNoContent)
}

func createEntry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filterName := vars["filterName"]
	entryName := r.FormValue("name")

	filter, ok := filters[filterName]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	filter.Add([]byte(entryName))
	w.WriteHeader(http.StatusCreated)
}

func retrieveEntry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filterName := vars["filterName"]
	entryName := vars["entryName"]

	filter, ok := filters[filterName]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	if !filter.Contains([]byte(entryName)) {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	w.WriteHeader(http.StatusOK)
	return
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
