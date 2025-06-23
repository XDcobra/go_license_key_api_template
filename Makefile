# Docker run commands
docker-start:
	docker-compose \
		-f "docker-compose.yml" \
		up \

docker-stop:
	docker-compose \
		-f "docker-compose.yml" \
		stop

docker-down:
	docker-compose \
		-f "docker-compose.yml" \
		down