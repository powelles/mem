# mem
Quick Memory Usage for Go

[![Build Status](https://travis-ci.org/powelles/mem.svg?branch=master)](https://travis-ci.org/powelles/mem)
[![GoDoc](https://godoc.org/github.com/powelles/mem?status.svg)](http://godoc.org/github.com/powelles/mem)
[![GoReportCard](https://img.shields.io/badge/go_report-A+-brightgreen.svg)](http://goreportcard.com/report/powelles/mem)

###Install:

    go get -u github.com/powelles/mem

###Usage:

    package main

    import (
        "fmt"

        "github.com/powelles/mem"
    )

    func main() {
        fmt.Println(mem.Allocated())
    }
