#!/bin/sh
go build main.go
go build -buildmode=plugin module.go

