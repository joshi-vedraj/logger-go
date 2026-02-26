package main

import (
	"fmt"

	"github.com/joshi-vedraj/logger-go/logger"
)

func main() {

	fmt.Println("**** LLD - Logger - Start ****")

	log := logger.GetLogger()

	fileWriter := logger.NewFileWriter("app.log")
	log.AddWriter(fileWriter)
	log.SetLevel(logger.DEBUG)

	log.Debug("debugging")
	log.Info("server started")
	log.Warn("Warning!")
	log.Error("db failed")

	fmt.Println("***** LLD - Logger - End *****")

}
