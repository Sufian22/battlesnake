package main

import (
	"context"
	"encoding/json"
	"flag"
	"os"
	"os/signal"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/etoshi/testingroom/battlesnake/pkg/config"
	"gitlab.com/etoshi/testingroom/battlesnake/pkg/server"
)

func main() {
	configPath := flag.String("config", "config.json", "Battlesnake server configuration")
	flag.Parse()

	logger := logrus.New()

	if configPath == nil || *configPath == "" {
		logger.Fatal("server configuration file not specified")
	}

	configFile, err := os.ReadFile(*configPath)
	if err != nil {
		logger.Fatal(err)
	}

	config := config.ServerConfig{}
	if err := json.Unmarshal(configFile, &config); err != nil {
		logger.Fatal(err)
	}

	logLevel, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		logger.Fatal(err)
	}

	logger.SetLevel(logLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		PadLevelText: true,
	})

	createBattleSnakeServerFunc := server.NewBattlesnakeServerFunc(logger)
	server, err := createBattleSnakeServerFunc(config)
	if err != nil {
		logger.Fatal(err)
	}

	go server.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	server.Shutdown(ctx)
	logger.Println("shutting down the server...")

	os.Exit(0)
}
