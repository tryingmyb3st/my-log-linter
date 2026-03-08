package example

import (
	"log/slog"

	"go.uber.org/zap"
)

const (
	token = "super secret"
	password
	apiKey
)

func main() {
	logger, _ := zap.NewProduction()

	//❌Неправильно
	logger.Info("Starting server on port 8080") // want "first letter should be in lowercase"
	slog.Error("Failed to connect to database") // want "first letter should be in lowercase"

	//✅Правильно
	logger.Info("starting server on port 8080")
	slog.Error("failed to connect to database")

	//❌Неправильно
	logger.Info("запуск сервера")                  // want "all logs should be in english"
	slog.Error("ошибка подключения к базе данных") // want "all logs should be in english"

	//✅Правильно
	logger.Info("starting server")
	slog.Error("failed to connect to database")

	//❌Неправильно
	logger.Info("server started!🚀")               // want "logs shouldn't contain special symbols"
	logger.Error("connection failed!!!")          // want "logs shouldn't contain special symbols"
	slog.Warn("warning: something went wrong...") // want "logs shouldn't contain special symbols"

	//✅Правильно
	logger.Info("server started")
	logger.Error("connection failed")
	slog.Warn("something went wrong")

	//❌Неправильно
	logger.Info("user password: " + password) // want "logs shouldn't contain credentials"
	logger.Debug("api_key=" + apiKey)         // want "logs shouldn't contain credentials"
	slog.Info("token: " + token)              // want "logs shouldn't contain credentials"

	//✅Правильно
	logger.Info("user authenticated successfully")
	logger.Debug("api request completed")
	slog.Info("token validated")
}
