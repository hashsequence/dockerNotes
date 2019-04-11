package main

import (
  	"database/sql"
    "fmt"
    "log"
      _ "github.com/lib/pq"
      "encoding/json"
)

const (
  host     = "localhost"
  port     =  5432
  user     = "postgres"
  password = "password"
  dbname   = "HashSequencePortfolio"
)

var Db *sql.DB
//each field must be capitalied to be exported
//https://stackoverflow.com/questions/28411394/golang-and-json-with-array-of-struct



func init() {
  var err error
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
  "password=%s dbname=%s sslmode=disable",
  host, port, user, password, dbname)

  fmt.Println("db init at port ", port)

  Db, err = sql.Open("postgres", psqlInfo)
  if err != nil {
    fmt.Println("db failed")
    log.Fatal(err)
  }
  fmt.Println("db success")
  return
}

func main() {
  rows, err := Db.Query("SELECT * from algos;")
  cols, _ := rows.Columns()
  store := []map[string]interface{}{}
  //fmt.Println(rows)
  for rows.Next() {
      columns := make([]interface{}, len(cols))
      columnPointers := make([]interface{}, len(cols))
      for i, _ := range columns {
          columnPointers[i] = &columns[i]
      }

      if err := rows.Scan(columnPointers...); err != nil {
          log.Fatal(err)
      }
      m := make(map[string]interface{})
      for i, colName := range cols {
          val := columnPointers[i].(*interface{})
          m[colName] = *val
      }
      store = append(store, m)
  }

  js, _ := json.MarshalIndent(store, "", "  ")
  fmt.Println(string(js))


  defer rows.Close()
  if err != nil {
    panic(err.Error())
  }
}
