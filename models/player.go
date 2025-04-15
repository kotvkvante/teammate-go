package models

type Player struct {
  ID int `json:"id"`
  PID int `json:"pid"`
  Nick string `json:"nick"`
  HP int `json:"hp"`
  Arm int `json:"arm"`
  Team_ID int `json:"team_id"`
}
