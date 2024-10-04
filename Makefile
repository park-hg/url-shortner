up-infra:
	docker compose -f compose/local.yaml up -d

down-infra:
	docker compose -f compose/local.yaml down

run:
	go run ./cmd/server/main.go

ddl:
	go run ./cmd/ddl/main.go
