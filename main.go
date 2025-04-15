package main

import (
  "os"
  "io"
  // "fmt"
  // "encoding/json"
  // "database/sql"

  "github.com/gin-gonic/gin"

  "test3/db"
  "test3/server"
)

func SetupLogOutput() {
  f, _ := os.Create("gin.log")
  gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
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
