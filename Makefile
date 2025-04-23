POSTGRESQL_URL="postgresql://postgres:postgres@localhost:15432/e_voting_master?sslmode=disable"

migrate-up:
	migrate -database ${POSTGRESQL_URL} -path db/migrations up

migrate-down:
	migrate -database ${POSTGRESQL_URL} -path db/migrations down 1

all:
	go run cmd/main.go

api:
	go run cmd/main.go api -c configs/config.yaml

scheduler:
	go run cmd/main.go scheduler -c configs/config.yaml

worker:
	go run cmd/main.go worker -c configs/config.yaml
