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

docker-build:
	docker-compose build
	docker push localhost:5000/api-gateway:latest

# Kubernetes locally
# Starting local dev helm for the first time
first-start:
	minikube start	# start minikube
	minikube addons enable ingress	# enable ingress for helm later
	kube-env # set minikube context


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

# Reloads
kube-reload:
	kubectl rollout restart deployment

helm-reload:
	helm upgrade --install gofiber-starter-stack ./charts/gofiber-starter-stack \
      -f ./charts/gofiber-starter-stack/values.yaml \
      -f ./charts/gofiber-starter-stack/secret-values.yaml

local-reg:
	kubectl port-forward --namespace kube-system service/registry 5000:80
	docker run --rm -it --network=host alpine ash -c "apk add socat && socat TCP-LISTEN:5000,reuseaddr,fork TCP:host.docker.internal:5000"