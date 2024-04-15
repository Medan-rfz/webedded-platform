
build-all:
	cd authorization && GOOS=linux GOARCH=amd64 make build
	cd users_management && GOOS=linux GOARCH=amd64 make build

run-all:
	docker-compose up --force-recreate --build
