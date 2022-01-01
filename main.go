package main

import (
	"github.com/adityapwr/banking-lib/logger"
	"github.com/adityapwr/go-banking/app"
)

func main() {
	logger.Info("Starting the application...")
	app.StartApp()
}
