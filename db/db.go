package db

import (
  "fmt"
  "log"
  "os"
  "strconv"
  "database/sql"


  "github.com/joho/godotenv"
  _ "github.com/lib/pq"
)

var Db *sql.DB

func ConnectDatabase()  {
   err := godotenv.Load()

   if err != nil {
      log.Println(err)
      return
   }

  host := os.Getenv("HOST")
  port, _ := strconv.Atoi(os.Getenv("PORT"))
  user := os.Getenv("USER")
  dbname := os.Getenv("DB_NAME")
  pass := os.Getenv("PASSWORD")

  conninfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
       host, port, user, dbname, pass)

  db, errSql := sql.Open("postgres", conninfo)
  if errSql != nil {
    log.Println(err)
    return
  } else {
    Db = db
    log.Println("Successfully connected to database!")
  }
}
