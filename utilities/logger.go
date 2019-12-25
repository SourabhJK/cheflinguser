package utilities

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/onrik/logrus/filename"
)

func GetLogger() {

	// open a file
	f, err := os.OpenFile("activity.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	log.New()

	//log.NewEntry(&l)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	log.SetOutput(f)

	filenameHook := filename.NewHook()
	filenameHook.Field = "source"
	log.AddHook(filenameHook)
}