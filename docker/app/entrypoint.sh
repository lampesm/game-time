#!/bin/bash

chown -R www-data:www-data /app

go env -w GOPROXY=https://goproxy.cn,https://gocenter.io,https://goproxy.io,direct

go mod download

go run main.go
