package main

import (
    "log"
		"net/http"
		"github.com/gorilla/mux"
		"encoding/json"
)

type Profile struct {
  Name    string
  Hobbies []string
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	// profile := Profile{"Alex", []string{"snowboarding", "programming"}}
	profileArr := []Profile { 
    Profile {
			Name: "Alex1", 
			Hobbies: []string{"snowboarding", "programming"},
    },
    Profile {
			Name: "Alex2", 
			Hobbies: []string{"snowboarding", "programming"},
    },
}
	json, err := json.Marshal(profileArr)
	if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
  w.Write(json)
}