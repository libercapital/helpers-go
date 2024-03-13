# DocGen

![Version](https://img.shields.io/badge/version-1.0.0-blue.svg?cacheSeconds=2592000) [![pipeline status](https://gitlab.com/bavatech/architecture/software/libs/go-modules/docgen/badges/main/pipeline.svg)](https://gitlab.com/bavatech/architecture/software/libs/go-modules/docgen/-/commits/main) [![coverage report](https://gitlab.com/bavatech/architecture/software/libs/go-modules/docgen/badges/main/coverage.svg)](https://gitlab.com/bavatech/architecture/software/libs/go-modules/docgen/-/commits/main)

Brazillian document CPF and CNPJ generator for GO.

Jump To:

- [DocGen](#docgen)
  - [Getting Started](#getting-started)
    - [Installing](#installing)
    - [Go Modules](#go-modules)
  - [Quick Examples](#quick-examples)

## Getting Started

### Installing

Use `go get` to retrieve the SDK to add it to your `GOPATH` workspace, or
project's Go module dependencies.

    go get github.com/libercapital/helpers-go/docgen

To update the module use `go get -u` to retrieve the latest version of the module.

    go get -u github.com/libercapital/helpers-go/docgen

### Go Modules

If you are using Go modules, your `go get` will default to the latest tagged
release version of the module. To get a specific release version of the module use
`@<tag>` in your `go get` command.

    go get github.com/libercapital/helpers-go/docgen@v1.0.0

To get the latest module repository change use `@latest`.

    go get github.com/libercapital/helpers-go/docgen@latest

## Quick Examples

This examples show how to generate random CPF and CNPJ.

Random functions need no argument and the seed ones takes an `int64` as argument and always generate the same value.
```
import (
    "github.com/libercapital/helpers-go/docgen"
)

docgen.RandomCNPJ()

docgen.RandomCPF()

docgen.SeedCNPJ(123)

docgen.SeedCPF(123)
```
