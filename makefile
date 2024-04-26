build_amd:
	GOOS=linux GOARCH=amd64 go build -o greet-server . 
docker_build_amd:
	docker buildx build --no-cache --platform linux/amd64  -t izaakdale/greet-server .
build:
	GOOS=linux go build -o greet-server . 
docker_build:
	docker build --no-cache -t izaakdale/greet-server .