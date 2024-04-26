#!/bin/bash

go run script/database/main.go

sleep 2

go run script/migration/main.go

tail -f /dev/null