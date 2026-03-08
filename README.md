# Линтер для проверки лог-записей

## Описание

Линтер для Go, совместимый с golangci-lint, который может анализировать лог-записи в коде и проверять их соответствие установленным правилам.

## Требования к линтеру

Линтер должен проверять следующие правила для лог-записей:

1. Лог-сообщения должны начинаться со строчной буквы

```go
//❌Неправильно
log.Info("Starting server on port 8080")
slog.Error("Failed to connect to database")

//✅Правильно
log.Info("starting server on port 8080")
slog.Error("failed to connect to database")
```

2. Лог-сообщения должны быть только на английском языке

```go
//❌Неправильно
log.Info("запуск сервера")
log.Error("ошибка подключения к базе данных")

//✅Правильно
log.Info("starting server")
log.Error("failed to connect to database")
```

3. Лог-сообщения не должны содержать спецсимволы или эмодзи

```go
//❌Неправильно
log.Info("server started!🚀")
log.Error("connection failed!!!")
log.Warn("warning: something went wrong...")

//✅Правильно
log.Info("server started")
log.Error("connection failed")
log.Warn("something went wrong")
```

4. Лог-сообщения не должны содержать потенциально чувствительные данные

```go
//❌Неправильно
log.Info("user password: " + password)
log.Debug("api_key=" + apiKey)
log.Info("token: " + token)

//✅Правильно
log.Info("user authenticated successfully")
log.Debug("api request completed")
log.Info("token validated")
```

## Сборка

```console
make build
```

## Запуск

```console
./bin/mylinter "path_to_file"
```

## Примеры использования:

```console

./bin/mylinter cmd/example/example.go 

C:\Users\User\my-log-linter\cmd\example\example.go:19:14: first letter should be in lowercase
C:\Users\User\my-log-linter\cmd\example\example.go:20:13: first letter should be in lowercase
C:\Users\User\my-log-linter\cmd\example\example.go:27:14: all logs should be in english
C:\Users\User\my-log-linter\cmd\example\example.go:28:13: all logs should be in english
C:\Users\User\my-log-linter\cmd\example\example.go:35:14: logs shouldn't contain special symbols
C:\Users\User\my-log-linter\cmd\example\example.go:36:15: logs shouldn't contain special symbols
C:\Users\User\my-log-linter\cmd\example\example.go:37:12: logs shouldn't contain special symbols
C:\Users\User\my-log-linter\cmd\example\example.go:45:14: logs shouldn't contain credentials
C:\Users\User\my-log-linter\cmd\example\example.go:46:15: logs shouldn't contain credentials
C:\Users\User\my-log-linter\cmd\example\example.go:47:12: logs shouldn't contain credentials

cd golangci-lint/ && make build && cd ../
./golangci-lint/golangci-lint.exe run cmd/example/example.go

example.go:19:14: myLogLinter: first letter should be in lowercase (myLinter)
        logger.Info("Starting server on port 8080") // want "first letter should be in lowercase"
                    ^
example.go:20:13: myLogLinter: first letter should be in lowercase (myLinter)
        slog.Error("Failed to connect to database") // want "first letter should be in lowercase"
                   ^
example.go:27:14: myLogLinter: all logs should be in english (myLinter)
        logger.Info("запуск сервера")                  // want "all logs should be in english"
                    ^
example.go:28:13: myLogLinter: all logs should be in english (myLinter)
        slog.Error("ошибка подключения к базе данных") // want "all logs should be in english"
                   ^
example.go:35:14: myLogLinter: logs shouldn't contain special symbols (myLinter)
        logger.Info("server started!🚀")               // want "logs shouldn't contain special symbols"
                    ^
example.go:36:15: myLogLinter: logs shouldn't contain special symbols (myLinter)
        logger.Error("connection failed!!!")          // want "logs shouldn't contain special symbols"
                     ^
example.go:37:12: myLogLinter: logs shouldn't contain special symbols (myLinter)
        slog.Warn("warning: something went wrong...") // want "logs shouldn't contain special symbols"
                  ^
example.go:45:14: myLogLinter: logs shouldn't contain credentials (myLinter)
        logger.Info("user password: " + password) // want "logs shouldn't contain credentials"
                    ^
example.go:46:15: myLogLinter: logs shouldn't contain credentials (myLinter)
        logger.Debug("api_key=" + apiKey)         // want "logs shouldn't contain credentials"
                     ^
example.go:47:12: myLogLinter: logs shouldn't contain credentials (myLinter)
        slog.Info("token: " + token)              // want "logs shouldn't contain credentials"
                  ^
10 issues:
* myLinter: 10

```