package main

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labring/sealos/service/devbox/api"
	tag "github.com/labring/sealos/service/devbox/pkg/registry"
)

func main() {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	tag.Init(user, password)
	if os.Getenv("GIN_MODE") != gin.DebugMode {
		gin.SetMode(gin.ReleaseMode)
	}
	slog.SetLogLoggerLevel(slog.LevelInfo)
	if os.Getenv("LOG_LEVEL") != "debug" {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	r := gin.Default()
	r.POST("/tag", api.Tag)
	if err := r.Run(":8092"); err != nil {
		slog.Error("Failed to start server", "Error", err)
		return
	}
}
