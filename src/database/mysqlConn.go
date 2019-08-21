package database

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

const DRIVER    = "mysql"
const USER_NAME = "root"
const PASSWORD  = ""
const DB_NAME   = "go_crud"

func Connect() (db * sql.DB){
    db, err := sql.Open(DRIVER, USER_NAME + ":" + PASSWORD + "@/" + DB_NAME)
    if err != nil{
      panic(err)
    }

    return db
}
