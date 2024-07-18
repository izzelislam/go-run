package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var baseUrl string = "http://localhost:8000"

type Student struct {
	ID string
	Name string
	Grade int
}

func fetchUser() ([]Student, error) {
	var err error
	var data []Student
	var client = &http.Client{}

	request, err := http.NewRequest("GET", baseUrl+"/users", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func Fetch() {
	usr, err := fetchUser()
	if err != nil {
		panic(err.Error())
		return
	}

	for _, v := range usr {
		fmt.Printf("ID: %s\t Name: %s\t Grade %d\n", v.ID, v.Name, v.Grade)
	}
}