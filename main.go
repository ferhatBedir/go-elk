package main

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"os"
	"time"
)

const logPath = "./logs/go.log"

var logger *zap.Logger

func main() {
	setupLog()
	r := gin.Default()
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))
	r.GET("/log", logHandler)
	_ = r.Run()

}

func setupLog() {
	file, _ := os.OpenFile(logPath, os.O_RDONLY|os.O_CREATE, 0666)
	log.SetOutput(file)
	c := zap.NewProductionConfig()
	c.OutputPaths = []string{"stdout", logPath}
	logger, _ = c.Build()

}

func logHandler(c *gin.Context) {
	logger.Warn("elk test", zap.String("elk", "test"))
	c.JSON(200, gin.H{
		"message": "Hello world",
	})

}
