package util

import "log"

func LogFatalIfErr(message string, err error) {
	if err != nil {
		log.Fatalf("%s, error: %+v", message, err)
	}

}
