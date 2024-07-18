package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	id string
	name string
	age int
	grade int
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "dev:dev@tcp(127.0.0.1:3306)/go_learn")

	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlQuery(){
	/* 
		connect to database
	*/
	db, err := connect();
	if err != nil {
		panic(err.Error())
		return
	}
	defer db.Close()

	var age = 27
	rows, err := db.Query("select id, name, grade from tb_student where age = ?", age)
	if err != nil {
		panic(err.Error())
		return
	}
	defer rows.Close()

	var result []student
	for rows.Next() {
		var each student
		var err = rows.Scan(&each.id, &each.name, &each.grade)
		if err != nil {
			panic(err.Error())
			return
		}
		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		panic(err.Error())
		return
	}

	for _, v := range result {
		fmt.Println(v.name)
	}

}

func main() {
	sqlQuery()
}