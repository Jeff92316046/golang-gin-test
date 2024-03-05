package api

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type Data_structure struct {
	Id     int    `json:"id"`
	Stock  string `json:"stock"`
	Date   string `json:"date"`
	Number string `json:"number"`
	People string `json:"people"`
	Share  string `json:"share"`
}

func GET_testdata(c *gin.Context) {

	data := connect_sql()
	fmt.Println(data[333].Date)
	j, err := json.Marshal(data)
	checkErr(err)

	c.Data(200, "application/json", j)
}

func connect_sql() []Data_structure {
	data := []Data_structure{}
	db, err := sql.Open("sqlite3", "./db.sqlite3")
	checkErr(err)
	defer db.Close()
	rows, err := db.Query("select * from data ")
	checkErr(err)
	for i := 0; rows.Next(); i++ {
		data = append(data, Data_structure{})
		var id int
		var stock string
		var date string
		var num string
		var people string
		var share string
		err := rows.Scan(&id, &stock, &date, &num, &people, &share)
		checkErr(err)
		data[i].Id, data[i].Stock, data[i].Date, data[i].Number, data[i].People, data[i].Share = id, stock, date, num, people, share

	}
	return data
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
