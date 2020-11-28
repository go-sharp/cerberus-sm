@echo off

gofmt -s -w .
gofmt -s -w ../
go run .\scripts\build.go
