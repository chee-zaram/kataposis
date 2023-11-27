package main

import (
	"fmt"
	"os"
	"time"

	"cheezaram.tech/kataposis"
)

func main() {
	var err error

	// Initialize the logger.
	log, err := kataposis.Init(
		kataposis.WithPGAddr("localhost:5432"),
		kataposis.WithPGDB("kataposis"),
		kataposis.WithPGUser("kataposis"),
		kataposis.WithPGPassword("kataposis"),
	)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	// Log a message. Timestamp must be called to store to the database.
	err = log.Msg("Log message").Level("info").RID("1234").Timestamp(time.Now())
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	t := time.Now()
	// You can query the database for log information.
	// This fetches all the logs with level equal to "info" and logged before
	// time.Now().
	result, err := log.Fetch("", "", "info", nil, &t)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	for _, res := range result {
		fmt.Printf("rid=%s, msg=%s, time=%s\n",
			res.GetRID(), res.GetMsg(), res.GetTimestamp(),
		)
	}
}
