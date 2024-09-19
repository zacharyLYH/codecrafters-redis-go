package main

import (
	"fmt"
	"os"
)

func handleGenericError(err error, message ...string) {
	defaultMessage := "generic error: "
	if len(message) > 0 {
		defaultMessage = message[0]
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s%v\n", defaultMessage, err)
		os.Exit(1)
	}
}
