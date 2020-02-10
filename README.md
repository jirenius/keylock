# Key Lock for Go

[![License MIT](https://img.shields.io/badge/License-MIT-blue.svg)](http://opensource.org/licenses/MIT)
[![Report Card](http://goreportcard.com/badge/github.com/jirenius/keylock)](http://goreportcard.com/report/jirenius/keylock)
[![Build Status](https://travis-ci.com/jirenius/keylock.svg?branch=master)](https://travis-ci.com/jirenius/keylock)
[![Reference](https://img.shields.io/static/v1?label=reference&message=go.dev&color=5673ae)](https://pkg.go.dev/github.com/jirenius/keylock)

A [Go](http://golang.org) package that provides mutex locking based on a key string.

## Installation

```bash
go get github.com/jirenius/keylock
```

## Usage

```go
    // Zero value of KeyLock is ready to use
    kl := &KeyLock{}

    // Lock key foo
    kl.Lock("foo")
    defer kl.Unlock("foo")

    // Read lock key bar
    kl.RLock("bar")
    defer kl.RUnlock("bar")
```
