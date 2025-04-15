package controllers

import (
  // "encoding/json"
  // "database/sql"

  "net/http"
  "github.com/gin-gonic/gin"

  "test3/db"
  "test3/models"
)

type TeamController struct {}

func (t TeamController) Get(ctx *gin.Context) {
  var teams []models.Team;

  rows, err := db.Db.Query("SELECT id, name FROM team")

  if err != nil {
    // err
  }
  defer rows.Close()

  for rows.Next() {
    var team models.Team
    err := rows.Scan(&team.ID, &team.Name)
    if err != nil {
      // err
    }
    teams = append(teams, team)
  }

  ctx.JSON(http.StatusOK, teams)
}

func (t TeamController) GetPlayers(ctx *gin.Context) {
  var nicks []string;

  team_id := ctx.Param("team_id")

  rows, err := db.Db.Query("SELECT nick FROM player WHERE team_id = $1", team_id)

  if err != nil {
    // err
  }
  defer rows.Close()

  for rows.Next() {
    var nick string
    err := rows.Scan(&nick)
    if err != nil {
      // err
    }
    nicks = append(nicks, nick)
  }

  ctx.JSON(http.StatusOK, nicks)
}
