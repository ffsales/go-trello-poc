package utils

import "log"

func ReturnError(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}
