package main

import (
	"github.com/ricardomaricato/banking/app"
	"github.com/ricardomaricato/banking/logger"
)

func main() {

	logger.Info("Starting the application...")
	app.Start()

}
