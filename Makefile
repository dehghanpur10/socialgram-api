SHELL := /bin/bash
.DEFAULT_GOAL := build

.EXPORT_ALL_VARIABLES:
PORT ?= 8000
DB_USER ?= mohammad
DB_PASSWORD ?= 173946285
DB_HOST ?= localhost
DB_PORT ?= 3306
DB_NAME ?= socialgram
DB_ENGINE ?= MYSQL
SECRET_KEY ?= super secret key
PAGE_SIZE ?= 20
build:
	go build -o main.exe main.go
	./main.exe

