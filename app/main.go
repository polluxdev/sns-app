package main

import (
	"github.com/polluxdev/sns-app/app/infrastructure"
)

func main() {
	logger := infrastructure.NewLogger()

	infrastructure.Load(logger)

	sqlHandler, err := infrastructure.NewSQLHandler()
	if err != nil {
		logger.LogError("%s", err)
	}

	infrastructure.Dispatch(logger, sqlHandler)
}
