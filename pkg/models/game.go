package models

import "github.com/sufian22/battlesnake/pkg/types"

type Game struct {
	ID      string       `json:"id"`
	Ruleset Ruleset      `json:"ruleset"`
	Timeout int32        `json:"timeout"`
	Source  types.Source `json:"source"`
}
