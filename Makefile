
build-all:
	cd users_management && GOOS=linux GOARCH=amd64 make build

run-all:
	docker-compose up --force-recreate --build
