package lib

import (
	"errors"
	"fmt"
	"strings"
)

func validate(input string) (bool, error){
	if strings.TrimSpace(input) == "" {
		return false, errors.New("canot be empty")
	}
	return true, nil
}

func catch (){
	if r := recover(); r != nil{
		fmt.Println("ada error", r)
	}else{
		fmt.Println("bagus")
	}
}

func Pnic(){
	defer catch()
	var name string
	println("masukan nama anda")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		println("hello", name)
	}else{
		panic(err.Error())
		println("end")
	}
}