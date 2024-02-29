help:	## Show this help.
	@fgrep -h "##" ${MAKEFILE_LIST} | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

build:	## Build application
	docker build -t jeanmolossi/rinha-backend-2024-q1:latest .

run:	## Run application
	docker-compose up -d
	docker logs -f api01 -n 30

reload:
	docker-compose down
	make build
	docker-compose up -d --force-recreate

run_dev:
	make build
	docker-compose up -d --force-recreate api01
	docker logs -f api01 -n 30

rebuild_db:
	docker rm db &>/dev/null
	docker-compose up db