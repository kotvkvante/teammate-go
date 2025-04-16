package controllers

import (
  "fmt"
  "encoding/json"
  "database/sql"

  "net/http"
  "github.com/gin-gonic/gin"

  "test3/db"
  "test3/models"
)

type PlayerController struct {}

func (p PlayerController) Get(ctx *gin.Context) {
  var players []models.Player;

  rows, err := db.Db.Query("SELECT id, pid, nick, hp, arm, team_id FROM player")

  if err != nil {
    fmt.Println("> ", err)
    return
  }
  defer rows.Close()

  for rows.Next() {
    var player models.Player
    err := rows.Scan(
      &player.ID, &player.PID, &player.Nick,
      &player.HP, &player.Arm, &player.Team_ID,
    )
    if err != nil {
      fmt.Println("> ", err)
      return
    }
    players = append(players, player)
  }

  ctx.JSON(http.StatusOK, players)
}

func (p PlayerController) GetByNick(ctx *gin.Context) {
  var player models.Player
  nick := ctx.Param("nick")

  row := db.Db.QueryRow("SELECT id, pid, nick, hp, arm, team_id FROM player WHERE nick = $1", nick)
  err := row.Scan(
    &player.ID, &player.PID, &player.Nick,
    &player.HP, &player.Arm, &player.Team_ID,
  )

  if err != nil {
    if err == sql.ErrNoRows {
      ctx.JSON(http.StatusNotFound, gin.H{"message": "player not found"})
      return
    }
    // err
    ctx.JSON(http.StatusNotFound, gin.H{"message": "?"})
    return
  }

  ctx.JSON(http.StatusOK, player)
}

func (p PlayerController) Add(ctx *gin.Context) {
  var player models.Player
  data, err := ctx.GetRawData()

  err = json.Unmarshal(data, &player)

  _, err = db.Db.Exec("INSERT INTO player (pid, nick, hp, arm, team_id) VALUES ($1, $2, $3, $4, $5)",
    player.PID,
    player.Nick,
    player.HP,
    player.Arm,
    player.Team_ID,
  )

  if err != nil {
    // fmt.Println(err)
    ctx.AbortWithStatusJSON(400, "Couldn't create the new player.")
  }

  ctx.JSON(http.StatusOK, "models.Player is successfully created.")
}
