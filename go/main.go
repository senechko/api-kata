package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type School struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var schools []School

func GetSchools(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(schools)
}

func GetSchool(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, school := range schools {
		if school.ID == id {
			json.NewEncoder(w).Encode(school)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func CreateSchool(w http.ResponseWriter, r *http.Request) {
	var school School
	_ = json.NewDecoder(r.Body).Decode(&school)

	schools = append(schools, school)

	json.NewEncoder(w).Encode("success")
}

func UpdateSchool(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i := range schools {
		if schools[i].ID == id {
			var s School
			_ = json.NewDecoder(r.Body).Decode(&s)
			s.ID = id
			schools[i] = s
			json.NewEncoder(w).Encode(schools[i])
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func ShowDocs(w http.ResponseWriter, r *http.Request) {
	r.Header.Add("Content-Type", "application/json")
	data, _ := ioutil.ReadFile("../swagger/swagger.json")
	w.Write(data)
}

func main() {

	for i := 1; i <= 3; i++ {
		schools = append(schools, School{ID: i, Name: fmt.Sprintf("School-%d", i)})
	}

	router := mux.NewRouter()
	router.HandleFunc("/schools", GetSchools).Methods("GET")
	router.HandleFunc("/schools/{id}", GetSchool).Methods("GET")
	router.HandleFunc("/schools", CreateSchool).Methods("POST")
	router.HandleFunc("/schools/{id}", UpdateSchool).Methods("PUT")
	router.HandleFunc("/docs/", ShowDocs).Methods("GET")
	http.ListenAndServe(":8080", router)
}
