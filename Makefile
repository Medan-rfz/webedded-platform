
build-all:
	cd users_management && make build

run-all:
	docker-compose up --force-recreate --build
