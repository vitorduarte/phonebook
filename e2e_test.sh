#!/bin/bash

# Up the minimal structure
docker-compose up --build -d mongodb app

# Run the tests
sleep 3
go test ./test/e2e/... -count=1 -timeout=3m -v

# Down the containers
docker-compose down -v
