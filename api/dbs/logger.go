package dbs

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/weekface/mgorus"
)

//Logger hook
func Logger() *logrus.Logger {
	logger := logrus.New()
	hooker, err := mgorus.NewHooker("localhost:27017", "login", "logs")
	if err == nil {
		logger.Hooks.Add(hooker)
	} else {
		fmt.Print(err)
	}

	return logger
}
