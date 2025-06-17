# Docker build commands
api-build:
	docker build \
		-t "api_gateway" \
		-f ".\api_gateway\\Dockerfile" \
	".\api_gateway"

# Docker run commands
docker-start:
	docker-compose \
		-f "docker-compose.yml" \
		up \

docker-stop:
	docker-compose \
		-f "C:\\Users\\Felix\\go\\src\\github.com\\XDcobra\\bmpapi\\docker-compose.yml" \
		stop

docker-down:
	docker-compose \
		-f "C:\\Users\\Felix\\go\\src\\github.com\\XDcobra\\bmpapi\\docker-compose.yml" \
		down