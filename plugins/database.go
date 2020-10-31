package plugins

import (
	"database/sql"
	"fmt"
	"time"
)

type DB struct {
	db *sql.DB
}

func initDataBase() *DB {
	fmt.Println("setting up database .... ")
	// db, err := sql.Open("database package name", "username:password@tcp(127.0.0.1:3306)/database name")
	db, err := sql.Open("mysql", "brian:xxxxxxx@tcp(127.0.0.1:3306)/coffeeshop")
	if err != nil {
		panic(err)
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(1000)
	db.SetMaxIdleConns(1000)

	return &DB{
		db: db,
	}
}

func (dBinst *DB) Query(queries string) (*sql.Rows, error) {
	rows, err := dBinst.db.Query(queries)

	return rows, err
}

// createUserTable := `CREATE TABLE IF NOT EXISTS User(
//     id int primary key auto_increment,
//     name varchar(128)not null UNIQUE,
//     password varchar(256)not null,
//     email varchar(128) UNIQUE,
//     isSuper BOOLEAN not null default 0
// )`

// insertRows, err = db.Query(createUserTable)

// if err != nil {
// 	panic(err.Error())
// }

// insertRows.Close()

// insertNewRow := `INSERT INTO
// 				User
// 				(name,password,email,isSuper)
// 				VALUES
// 				('brian','Wchh!983','brianwchh@gmail.com',true)`

// insertRows, err = db.Query(insertNewRow)

// if err != nil {
// 	panic(err.Error())
// }

// insertRows.Close()
