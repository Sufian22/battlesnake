package config

import "github.com/BattlesnakeOfficial/rules/cli/commands"

type ServerConfig struct {
	Port        string                `json:"port"`
	Environment string                `json:"env"`
	LogLevel    string                `json:"logLevel"`
	Info        commands.PingResponse `json:"info"`
}
