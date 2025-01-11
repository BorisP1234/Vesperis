package main

import (
	"github.com/team-vesperis/vesperis/proxy/database"
	"github.com/team-vesperis/vesperis/proxy/logger"
	"go.minekube.com/gate/cmd/gate"
)

func main() {
	logger.CreateLogger()
	logger := logger.GetLogger()
	logger.Info("Starting Vesperis Proxy...")
	database.InitializeDatabases()
	gate.Execute()
}
