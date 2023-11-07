run-with-build:
	docker-compose -f ./docker-compose.yml --env-file=.env up --build --abort-on-container-exit --exit-code-from app

run:
	docker-compose -f ./docker-compose.yml --env-file=.env up --abort-on-container-exit --exit-code-from app

stop:
	docker-compose down --remove-orphans

build:
	go build -mod=vendor -o app.out -v ./cmd/server/main.go

build-docker-server: # запуск обычного http servera
	docker build -f ./Dockerfile . --tag rsoi1