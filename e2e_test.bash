docker-compose down -v
docker-compose up --build -d mongodb app

go test ./test/e2e/... -count=1 -timeout=3m
