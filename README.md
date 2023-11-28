<div align="center">
  <strong>A Go logging library with database integration</strong>

<h1>Kataposis</h1>
</div>

[![Workflow](https://github.com/chee-zaram/kataposis/actions/workflows/go.yml/badge.svg)][workflow]
[![Go Report](https://goreportcard.com/badge/github.com/chee-zaram/kataposis)][report]
![Last Commit](https://img.shields.io/github/last-commit/chee-zaram/kataposis)
![Contributors](https://img.shields.io/github/contributors/chee-zaram/kataposis)

---

## Table of Contents

- [Introduction](#introduction)
- [Installation](#installation)
- [Usage](#usage)
  - [Initialization](#initialization)
  - [Logging Messages](#logging-messages)
  - [Querying Log Entries](#querying-log-entries)
  - [Database Schema](#database-schema)
- [Author](#author)
- [Contributing](#contributing)
- [Licensing](#licensing)

## Introduction

**Kataposis** (or `ktpss`) is a logging library for Go that provides a simple
but effective logging mechanism, seemlessly saving entries to a database. It
provides a fluent API for constructing and storing log entries in the database,
and it also supports fetching log entries based on various criteria.

At the moment, supported databases include Postgres. Kataposis currently
supports the basic log parameters such as message, log level, timestamp, and
resource ID; with a design that allows for easy support for more parameters as
needed in the future.

## Installation

First step towards using Kataposis in your Go application is to add the module
to your Go program:

```sh
go get -u cheezaram.tech/kataposis
```

## Usage

#### Initialization

To initialize Kataposis and set up the database connection, use the `Init`
function along with optional configuration options:

```go
package main

import (
	"fmt"
	"os"
	"time"

	"cheezaram.tech/kataposis"
)

func main() {
	// Initialize the logger.
	log, err := kataposis.Init(
		kataposis.WithPGAddr("localhost:5432"),
		kataposis.WithPGDB("kataposis"),
		kataposis.WithPGUser("kataposis"),
		kataposis.WithPGPassword("kataposis"),
	)

	if err != nil {
		// Handle error
	}
}
```

#### Logging Messages

Use the fluent API provided by Kataposis to create log entries and save them to
the database:

```go
func main() {
	// Previous initialization code.

	// Log a message.
	err := log.Msg("Log message").RID("1234").Level("info").Timestamp(time.Now())
	if err != nil {
		// Handle error
	}
}
```

#### Querying Log Entries

You can fetch log entries from the database based on specific criteria using the
`Fetch` method:

```go
func main() {
	// Previous initialization code.

	afterTime := time.Now()
	entries, err := log.Fetch(
		"message",  // Message
		"1234",     // Resource ID
		"",         // Log level
		nil,        // Before timestamp
		&afterTime, // After timestamp
	)

	if err != nil {
		// Handle error
	}

	// Process fetched log entries
	for _, entry := range entries {
	    // Do something with each log entry
	    fmt.Printf("Message: %s, Timestamp: %s\n", entry.GetMsg(), entry.GetTimestamp())
	}
}
```

You can retrieve the value of each field using the methods of the `log.Entry`
struct prefixed with `Get`.

> NB: If any argument is set to its default value during quering, the argument
> will not be included in the query.

#### Database Schema

Kataposis creates a table named **logs** in the specified database with the
following schema:

```pg
CREATE TABLE IF NOT EXISTS logs (
    id SERIAL PRIMARY KEY,
    rid TEXT,
    message TEXT,
    level TEXT,
    timestamp TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

## Author

This library was written by [Chee-zaram Okeke](https://cheezaram.tech/) with
love. Feel free to connect.

## Contributing

If you find any bugs, feel free to
[open an issue](https://github.com/chee-zaram/kataposis/issues). Pull request
for new features or bug fixes are welcome.

After commiting your changes, run the
[utils/contributors.sh](./utils/contributors.sh) script to add yourself to the
contributors list.

See [CONTRIBUTORS](./CONTRIBUTORS) for details of all contributors.

## Licensing

See [LICENSE](./LICENSE) for details.

[workflow]: https://github.com/chee-zaram/kataposis/actions/workflows/go.yml?query=branch%3Amain+event%3Apush
[report]: https://goreportcard.com/report/github.com/chee-zaram/kataposis
