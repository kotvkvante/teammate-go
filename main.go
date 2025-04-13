package main

import (
  "os"
  "io"
  "fmt"
  "encoding/json"
  "database/sql"

  "net/http"
  "github.com/gin-gonic/gin"

  "test3/db"
)

type player struct {
  ID int `json:"id"`
  PID int `json:"pid"`
  Nick string `json:"nick"`
  HP int `json:"hp"`
  Arm int `json:"arm"`
}

func getPlayers(c *gin.Context) {
  var players []player;

  rows, err := db.Db.Query("SELECT id, pid, nick, hp, arm FROM player")

  if err != nil {
    // err
  }
  defer rows.Close()

  for rows.Next() {
    var pl player
    err := rows.Scan(&pl.ID, &pl.PID, &pl.Nick, &pl.HP, &pl.Arm)
    if err != nil {
      // err
    }
    players = append(players, pl)
  }

  c.JSON(http.StatusOK, players)
}

func getPlayerByNick(c *gin.Context) {
  var pl player
  nick := c.Param("nick")

  row := db.Db.QueryRow("SELECT id, pid, nick, hp, arm FROM player WHERE nick = $1", nick)
  err := row.Scan(&pl.ID, &pl.PID, &pl.Nick, &pl.HP, &pl.Arm)

  if err != nil {
    if err == sql.ErrNoRows {
      c.JSON(http.StatusNotFound, gin.H{"message": "player not found"})
      return
    }
    // err
    c.JSON(http.StatusNotFound, gin.H{"message": "?"})
    return
  }

  c.JSON(http.StatusOK, pl)
}

func addPlayer(c *gin.Context) {
  newplayer := player {}
  data, err := c.GetRawData()

  err = json.Unmarshal(data, &newplayer)

  _, err = db.Db.Exec("INSERT INTO player (pid, nick, hp, arm) VALUES ($1, $2, $3, $4)",
    newplayer.PID,
    newplayer.Nick,
    newplayer.HP,
    newplayer.Arm,
  )

  if err != nil {
    fmt.Println(err)
    c.AbortWithStatusJSON(400, "Couldn't create the new player.")
  } else {
    c.JSON(http.StatusOK, "Player is successfully created.")
  }
}

func setupLogOutput() {
  f, _ := os.Create("gin.log")
  gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
  setupLogOutput()
  router := gin.Default()

  db.ConnectDatabase()

  router.GET("/players", getPlayers)
  router.GET("/players/:nick", getPlayerByNick)

  router.POST("/add", addPlayer)

  router.Run(":8081")
}
