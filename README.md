# Organism

![Lint](https://github.com/citilinkru/organism/actions/workflows/golangci-lint.yml/badge.svg?branch=master)
![Tests](https://github.com/citilinkru/organism/actions/workflows/test.yml/badge.svg?branch=master)
[![codecov](https://codecov.io/gh/citilinkru/organism/branch/master/graph/badge.svg)](https://codecov.io/gh/citilinkru/organism)
[![Go Report Card](https://goreportcard.com/badge/github.com/citilinkru/organism)](https://goreportcard.com/report/github.com/citilinkru/organism)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/citilinkru/organism/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/citilinkru/organism?status.svg)](https://godoc.org/github.com/citilinkru/organism)
[![Release](https://img.shields.io/github/release/citilinkru/organism.svg?style=flat-square)](https://github.com/citilinkru/organism/releases/latest)

Abstraction for liveness and readiness probes of your app

# Description
We can describe each app as Organism, that consists of core and Limbs. Each Limb describes important part of your app, 
without which app can't work properly. If each Limb is ready to work, then the whole organism ready to work too 
(readiness probe). If at least one Limb is dead, then the whole organism is partially dead too.

# Example
Let's take simple http handlers funcs, to answer on readiness and liveness probes

```go
package main

import (
	"github.com/citilinkru/organism"
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

func main() {
	o := organism.New()
	limb1 := o.GrowLimb()
	go func() {
		defer limb1.Die()
		limb1.Ready()
		err := DoSmtValuable()
		if err != nil {
			log.Println("something wrong with 1: ", err)
		}
	}()

	limb2 := o.GrowLimb()
	go func() {
		defer limb2.Die()
		limb2.Ready()
		err := DoAnotherValuable()
		if err != nil {
			log.Println("something wrong with 2: ", err)
		}
	}()
	
	o.Ready()
	// ...
	
	r := mux.NewRouter()
	r.HandleFunc("/healty", ReadinessHandler(o))
	r.HandleFunc("/healtz", LivenessHandler(o))
	http.Handle("/", r)
	
	// ...
}

func LivenessHandler(o *organism.Organism) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if !o.IsAlive() {
			return
		}

		_, err := writer.Write([]byte("OK"))
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func ReadinessHandler(o *organism.Organism) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if !o.IsReady() {
			return
		}

		_, err := writer.Write([]byte("OK"))
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		}
	}
}
```

Testing
-----------
Unit-tests:
```bash
go test -v -race ./...
```

Run linter:
```bash
go mod vendor \
  && docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.40 golangci-lint run -v \
  && rm -R vendor
```

CONTRIBUTE
-----------
* write code
* run `go fmt ./...`
* run all linters and tests (see above)
* create a PR describing the changes

LICENSE
-----------
MIT

AUTHOR
-----------
Nikita Sapogov <amstaffix@gmail.com>