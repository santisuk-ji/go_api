package main

import (
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "database/sql"
  "log"
)

func initDB() (*sql.DB, error){
  db,err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/YFC_proj")
  if err != nil {
    fmt.Println("err1")
    log.Println(err)
  }
  // defer db.Close()
  return db,err
}

func prepareStatement(db *sql.DB,q string)(*sql.Stmt, error){
  stmt,err := db.Prepare(q)
  if err != nil {
    log.Fatal(err)
  }

  return stmt,err
}
