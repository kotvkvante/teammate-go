package main

import (
  "os"
  "io"
  "log"

  "github.com/gin-gonic/gin"

  "test3/db"
  "test3/server"
)

func SetupLogOutput() {
  f, _ := os.Create("gin.log")
  gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
  f, err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
  if err != nil {
      log.Fatalf("error opening file: %v", err)
  }
  defer f.Close()

  log.SetOutput(f)

  SetupLogOutput()
  db.ConnectDatabase()
  server.Init()

  // router := gin.Default()
  // router.GET("/players", getPlayers)
  // router.GET("/players/:nick", getPlayerByNick)
  //
  // router.POST("/add", addPlayer)

  // router.Run(":8081")
}
