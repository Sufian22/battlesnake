package config

import (
	"regexp"
)

type ServerConfig struct {
	Port        string          `json:"port"`
	Environment string          `json:"env"`
	LogLevel    string          `json:"logLevel"`
	Info        BattlesnakeInfo `json:"info"`
}

type BattlesnakeInfo struct {
	APIVersion string `json:"apiversion"`
	Author     string `json:"author"`
	Color      Color  `json:"color"`
	Head       string `json:"head"`
	Tail       string `json:"tail"`
	Version    string `json:"version"`
}

type Color string

const validColorRegexp = `(#\d{6})`

func (c *Color) IsValid() bool {
	ok, err := regexp.MatchString(validColorRegexp, string(*c))
	if err != nil || !ok {
		return false
	}

	return true
}
