package lib

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	Fulname string `json:"Name"`
	Age int
}

func EncodeDecode() {

	// ============== decode json to struct ==================
	var jsonString string = `{"Name": "john", "Age": 20}`
	var jsonData = []byte(jsonString)

	var singleUser User 
	err := json.Unmarshal(jsonData, &singleUser)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(singleUser.Fulname)
	fmt.Println(singleUser.Age)


	// ============== descode to map[string]interface{} =============
	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)

	println(reflect.ValueOf(data1["Name"]).String())
	println(reflect.ValueOf(data1["Age"]).Float())

	// ====== array json to array object

	var users = `[
		{"Name": "john wick", "Age": 27},
    {"Name": "ethan hunt", "Age": 32}
	]`

	var datas []User
	err_ := json.Unmarshal([]byte(users), &datas)

	if err_ != nil {
		panic(err_.Error())
	}

	println(datas[0].Fulname)
	println(datas[1].Fulname)

}