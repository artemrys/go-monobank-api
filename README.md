# Golang bindings for the Monobank API

[![GoDoc](https://godoc.org/github.com/artemrys/go-monobank-api?status.svg)](https://godoc.org/github.com/artemrys/go-monobank-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/artemrys/go-monobank-api)](https://goreportcard.com/report/github.com/artemrys/go-monobank-api)
[![Travis](https://travis-ci.org/artemrys/go-monobank-api.svg?branch=master)](https://travis-ci.org/artemrys/go-monobank-api)

All methods are fairly self explanatory, and reading the godoc page should
explain everything. If something isn't clear, open an issue or submit
a pull request.

The scope of this project is just to provide a wrapper around the API
without any additional features. There are other projects for creating
something with plugins and command handlers without having to design
all that yourself.

### Example

Please take a look at `examples/example.go`.

Run it using:
```bash
cd examples/
go run example.go -token <your-token>
```

### Links

- [Token](https://api.monobank.ua/)
- [Docs](https://api.monobank.ua/docs/)
