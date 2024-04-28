#!/bin/bash

rabbitmq-server -detached

sleep 5

rabbitmqctl import_definitions ./api-definitions.json

tail -f /dev/null