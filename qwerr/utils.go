package qwerr

import (
	"log"
)

func AbortWhen(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
