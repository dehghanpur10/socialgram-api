SHELL := /bin/bash
.DEFAULT_GOAL := build

.EXPORT_ALL_VARIABLES:
SERVER_PORT ?= :8000

build:
	go build -o main.exe main.go
	./main.exe


