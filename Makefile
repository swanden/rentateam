include .env

init: compose-build compose-up

compose-build:
	docker-compose build

compose-up:
	docker-compose up

compose-up:
	docker-compose down

compose-postgres:
	docker-compose up postgres

migrations-create:
	docker-compose run --rm migrate create -ext sql -dir rentateam/migrations -seq posts_table

migrations-up:
	docker-compose run --rm migrate -database postgres://user:password@localhost:5432/app?sslmode=disable -path rentateam/migrations up

migrations-down:
	docker-compose run --rm migrate -database ${PG_DSN} -path rentateam/migrations down

swagger-init:
	swag init -g ./internal/controller/http/v1/router.go

proto-generate:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/grpcpb/post.proto
