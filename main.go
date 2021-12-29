package main

import (
	"github.com/adityapwr/go-banking/app"
	"github.com/adityapwr/go-banking/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.StartApp()
}
