package main

import (
	"github.com/ashishjuyal/banking-lib/logger"
	"github.com/ricardomaricato/banking/app"
)

func main() {

	logger.Info("Starting the application...")
	app.Start()

}
