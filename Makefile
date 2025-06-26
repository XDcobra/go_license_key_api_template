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


# Kubernetes locally
# Set Docker to use Minikube's environment
kube-env:
	powershell -Command "minikube -p minikube docker-env | Invoke-Expression"

# apply k8s scripts
kube-apply:
	kubectl apply -f ./k8s -R

# Build docker compose image for minikube
kube-build:
	kube-env
	docker-compose build

# Reload all pods
kube-reload:
	kubectl rollout restart deployment
