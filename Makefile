
build-all:
	cd users_management && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 make build

run-all:
	docker-compose up --force-recreate --build
