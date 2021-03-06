include .env

init:
	docker-compose up --build

up:
	docker-compose up

down:
	docker-compose down

up-pg:
	docker-compose up postgres

migrations-create:
	docker-compose run --rm migrate create -ext sql -dir rentateam/migrations -seq posts_table

migrations-up:
	docker-compose run --rm migrate -database ${PG_DSN} -path rentateam/migrations up

migrations-down:
	docker-compose run --rm migrate -database ${PG_DSN} -path rentateam/migrations down

swagger-init:
	swag init -g ./internal/controller/http/v1/router.go

proto-generate:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/grpcpb/post.proto
