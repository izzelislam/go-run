package lib

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "index")
}

func Miniweb() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello")
	})

	http.HandleFunc("/index", index)

	http.HandleFunc("/greeting", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"Name": "John",
			"Message" : "Selamat datang",
		}

		t, err := template.ParseFiles("./pages/template.html")
		
		if err != nil {
			panic(err)
		}

		t.Execute(w, data)

	})

	println("server start on http://localhost:8000")
	http.ListenAndServe(":8000", nil);
}