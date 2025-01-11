package main

import (
	"github.com/vesperismc/vesperis/proxy/logger"
	"go.minekube.com/gate/cmd/gate"
)

func main() {
	logger.CreateLogger()
	logger := logger.GetLogger()
	logger.Info("Starting Vesperis Proxy...")
	gate.Execute()
}
