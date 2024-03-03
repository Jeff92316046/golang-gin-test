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
	number int
	people int
	share  int
}

func GET_testdata(data []Data_structure) {
	connect_sql(data)
}

func connect_sql(data []Data_structure) {
	db, err := sql.Open("sqlite3", "./db.sqlite3")
	checkErr(err)
	defer db.Close()
	rows, err := db.Query("select * from data")
	checkErr(err)
	for i := 0; rows.Next(); i++ {
		err := rows.Scan(data[i].id, data[i].stock, data[i].date, data[i].number, data[i].people, data[i].share)
		checkErr(err)
		fmt.Printf("%d", data[i].id)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
