SHELL := /bin/bash
.DEFAULT_GOAL := build

.EXPORT_ALL_VARIABLES:
SERVER_PORT ?= :8000
DB_USER ?= mohammad
DB_PASSWORD ?= 173946285
DB_HOST ?= localhost
DB_PORT ?= 3306
DB_NAME ?= socialgram

build:
	go build -o main.exe main.go
	./main.exe


