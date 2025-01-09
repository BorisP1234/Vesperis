package database

import (
	"github.com/go-logr/logr"
	c "github.com/spf13/viper"
)

func CreateConfig(logger *logr.Logger) {
	c.SetConfigName("network")
	c.SetConfigType("yaml")
	c.AddConfigPath(".")

	err := c.ReadInConfig()
	if err != nil {
		logger.Error(err, "Error loading the Database Config.")
		panic(err)
	}

	logger.Info("Loaded the Database Config.")
}

func GetMySQLUrl() string {
	return c.GetString("databases.mysql.username") +
		":" +
		c.GetString("databases.mysql.password") +
		"@(" +
		c.GetString("databases.mysql.host") +
		":" +
		c.GetString("databases.mysql.port") +
		")/" +
		c.GetString("databases.mysql.database")
}

func GetRedisUrl() string {
	host := c.GetString("databases.redis.host")
	port := c.GetString("databases.redis.port")
	database := c.GetString("databases.redis.database")
	username := c.GetString("databases.redis.username")
	password := c.GetString("databases.redis.password")

	if username != "" && password != "" {
		return "redis://" + username + ":" + password + "@" + host + ":" + port + "/" + database
	}

	return "redis://" + host + ":" + port + "/" + database
}
