SCRIPT_DIR=./scripts

.PHONY: init
init:
	${SCRIPT_DIR}/init.sh

.PHONY: build
build:
	${SCRIPT_DIR}/build.sh

.PHONY: migrate
migrate:
	${SCRIPT_DIR}/migrate.sh

.PHONY: docker
docker:
	${SCRIPT_DIR}/docker.sh
