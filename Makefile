.PHONY: run-pub run-sub run-tools docker-compose-up docker-compose-down

run-pub:
	go run ./pub/main.go $(topic)

run-sub:
	go run ./sub/main.go

run-tools:
	go run ./tools/main.go

docker-compose-up:
	docker compose up -d

docker-compose-down:
	docker compose down


