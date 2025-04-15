package server

import (
  "github.com/gin-gonic/gin"

  "test3/config"
  "test3/controllers"
)

func Init() {
  config.Init("config/config.yaml")

  router := gin.Default()

  player := new(controllers.PlayerController)
  player_group := router.Group("players")
  player_group.GET("/", player.Get)
  player_group.GET("/:nick", player.GetByNick)
  player_group.POST("/", player.Add)

  team := new(controllers.TeamController)
  team_group := router.Group("teams")
  team_group.GET("/", team.Get)
  team_group.GET("/:team_id", team.GetPlayers)

  router.Run(config.AppConfig.Addr)
}
