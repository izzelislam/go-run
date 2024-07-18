package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// type Student struct {
// 	ID string
// 	Name string
// 	Grade int
// }

var data = []Student{ 
	{"E001", "ethan", 21},
	{"W001", "wick", 22},
	{"B001", "bourne", 23},
	{"B002", "bond", 23},
}

func users(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "aplication.json")

	if r.Method == "GET" {
		res, err := json.Marshal(data)
		
		if err != nil {
			http.Error(w, "", http.StatusBadRequest)
		}
		w.Write(res)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func user(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var id = r.FormValue("id")

		for _, v := range data {
			if v.ID == id {
				res, err := json.Marshal(v)
				if err != nil {
					http.Error(w, "", http.StatusBadRequest)
					return
				}
				w.Write(res)
				return
			}
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		http.Error(w, "", http.StatusBadRequest)
	}
}

func ResApi() {

	http.HandleFunc("/users", users)
	http.HandleFunc("/user", user)

	

	fmt.Println("server runnig on http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}