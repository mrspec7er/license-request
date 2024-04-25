#!/bin/bash

go run script/database/main.go

sleep 2

go run script/migration/main.go

sleep 2

fresh

tail -f /dev/null