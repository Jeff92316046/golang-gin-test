package api

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Data_structure struct {
	id     int
	stock  string
	date   string
	number string
	people string
	share  string
}

func GET_testdata() {
	connect_sql()
}

func connect_sql() {
	data := []Data_structure{}
	db, err := sql.Open("sqlite3", "./db.sqlite3")
	checkErr(err)
	defer db.Close()
	rows, err := db.Query("select * from data ")
	checkErr(err)
	for i := 0; rows.Next(); i++ {
		data = append(data, Data_structure{})
		err := rows.Scan(&(data[i].id), &(data[i].stock), &(data[i].date), &(data[i].number), &(data[i].people), &(data[i].share))
		checkErr(err)
		fmt.Printf("%d %s %s %2s %9s %11s\n", data[i].id, data[i].stock, data[i].date, data[i].number, data[i].people, data[i].share)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
