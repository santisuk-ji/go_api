package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
    "strings"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    http.HandleFunc("/v1/psn/", GetPsnById)

    log.Fatal(http.ListenAndServe(":1234", nil))

    fmt.Println("server started on port :1234")
}

func GetPsnById(w http.ResponseWriter, r *http.Request) {
  psnid := strings.TrimPrefix(r.URL.Path, "/v1/psn/")
  // keys, _ := r.URL.Query()["key"]
  // fmt.Println(keys[0])
  db,err := initDB();
  defer db.Close()
  q := "select * from psn where id = ?"
  stmt,err := prepareStatement(db,q)
  defer stmt.Close()

  rows, err := stmt.Query(psnid)
  if err != nil {
    log.Fatal(err)
  }

  var id,fname,lname string
  for rows.Next() {
    err = rows.Scan(&id, &fname, &lname)
    if err != nil {
      fmt.Println("err4")
      log.Fatal(err)
    }
    fmt.Println("data : "+id+","+fname+","+lname)
  }
    fmt.Fprintf(w, "Test, %q", html.EscapeString(r.URL.Path))
}
