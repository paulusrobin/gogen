help: ## This help dialog
help h:
	@IFS=$$'\n' ; \
	help_lines=(`fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##/:/'`); \
	printf "%-30s %s\n" "target" "help" ; \
	printf "%-30s %s\n" "------" "----" ; \
	for help_line in $${help_lines[@]}; do \
		IFS=$$':' ; \
		help_split=($$help_line) ; \
		help_command=`echo $${help_split[0]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
		help_info=`echo $${help_split[2]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
		printf '\033[36m'; \
		printf "%-30s %s" $$help_command ; \
		printf '\033[0m'; \
		printf "%s\n" $$help_info; \
	done

#===================#
#== Env Variables ==#
#===================#
DOCKER_COMPOSE_FILE ?= docker-compose.yaml
RUN_IN_DOCKER ?= docker-compose exec -T gogen-project
RUN_TEST_IN_DOCKER ?= docker-compose exec -T -e APP_ENV=test gogen-project

#===============#
#== App Build ==#
#===============#

build-native: RUN_IN_DOCKER=
build-native: build

build: go-private-in-docker
	# Build APP Binary
	@echo "==========================="
	@echo "Building APP binary"
	@echo "==========================="
	${RUN_IN_DOCKER} go mod tidy && go mod vendor
	${RUN_IN_DOCKER} sh -c "go build -mod=vendor -o bin/gogen-project cmd/main.go"

run-subscriber:
	${RUN_IN_DOCKER} bin/gogen-project subscriber

run-http:
	${RUN_IN_DOCKER} bin/gogen-project http

run-grpc:
	${RUN_IN_DOCKER} bin/gogen-project grpc

clean: # clean executables
	rm -rf bin/*

#=======================#
#== ENVIRONMENT SETUP ==#
#=======================#

create-env-file:
ifeq (,$(wildcard .env))
	cp .env.sample .env
endif

remove-env-file:
ifneq (,$(wildcard .env))
	rm .env
endif

go-private:
	go env -w GOPRIVATE=github.com/paulusrobin

go-private-in-docker:
	${RUN_IN_DOCKER} go env -w GOPRIVATE=github.com/paulusrobin

docker-start:
	@echo "=========================="
	@echo "Starting Docker Containers"
	@echo "=========================="
	docker-compose -f ${DOCKER_COMPOSE_FILE} up -d --build --remove-orphans
	docker-compose -f ${DOCKER_COMPOSE_FILE} ps

docker-stop:
	@echo "=========================="
	@echo "Stopping Docker Containers"
	@echo "=========================="
	docker-compose -f ${DOCKER_COMPOSE_FILE} stop
	docker-compose -f ${DOCKER_COMPOSE_FILE} ps

docker-clean: docker-stop
	@echo "=========================="
	@echo "Removing Docker Containers"
	@echo "=========================="
	docker-compose -f ${DOCKER_COMPOSE_FILE} rm -v -f

docker-restart: docker-stop docker-start

environment: ## The only command needed to start a working environment
environment: remove-env-file docker-restart create-env-file go-private-in-docker build-native

environment-clean: ## The only command needed to clean the environment
environment-clean: docker-clean clean


#====================#
#== QUALITY CHECKS ==#
#====================#

static-analysis: ## Run all enabled linters
	@echo "======================="
	@echo "Running static analysis"
	@echo "======================="
	${RUN_IN_DOCKER} golangci-lint run
