package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// функция открытия БД
func OpenDatabase(path string) *sql.DB {
	dbb, err := sql.Open("sqlite3", path)
	db = dbb
	if err != nil {
		panic(err)
	}
	sql_request := `CREATE TABLE IF NOT EXISTS "archive" (
		"number"	INTEGER UNIQUE,
		"date"	TEXT,
		"for_whom"	TEXT,
		"from_whom"	TEXT,
		"for_what"	TEXT
	)`
	_, err = db.Exec(sql_request)
	if err != nil {
		panic(err)
	}
	return db
}

// функция закрытия БД
func CloseDatabase(db *sql.DB) {
	db.Close()
}

// функция добавления данных в БД
func InsertData(table_name string, rows_names string, values string) {
	db := OpenDatabase(configdata.Database_path + "database_" + strconv.Itoa(time.Now().Year()) + ".db")
	defer CloseDatabase(db)
	sql_request := "insert into " + table_name + " " + rows_names + " values " + values
	_, err := db.Exec(sql_request)
	if err != nil {
		panic(err)
	}
}

// функция получения данных с БД
func GetData() string {
	var data string
	db := OpenDatabase(configdata.Database_path + "database_" + strconv.Itoa(time.Now().Year()) + ".db")
	defer CloseDatabase(db)
	rows, _ := db.Query("select max(number) from archive")
	for rows.Next() {
		err := rows.Scan(&data)
		if err != nil {
			fmt.Println(err)
		}
	}
	return data
}

// функция, загружающая всю БД в структуру
// для дальнейшей загрузки на сайт в таблицу
func Show_database() DataFromDB {
	data := DataFromDB{}
	db := OpenDatabase(configdata.Database_path + "database_" + strconv.Itoa(time.Now().Year()) + ".db")
	defer CloseDatabase(db)
	rows, _ := db.Query("select * from archive")
	for rows.Next() {
		p := DataFromDB_view{}
		err := rows.Scan(
			&p.Number_of_document,
			&p.Date,
			&p.For_whom,
			&p.From_whom,
			&p.For_what)
		if err != nil {
			fmt.Println(err)
			continue
		}
		data.Data = append(data.Data, p)
	}
	return data
}

// функция удаления последней строки в БД
func DeleteLastRow() {
	db := OpenDatabase(configdata.Database_path + "database_" + strconv.Itoa(time.Now().Year()) + ".db")
	defer CloseDatabase(db)
	sql_request := "DELETE from archive WHERE number = (SELECT MAX(number) FROM archive)"
	_, err := db.Exec(sql_request)
	if err != nil {
		panic(err)
	}
}