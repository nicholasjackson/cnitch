# docker-machine guid for docker group
DOCKER_GROUP=100

test:
	go test -v -race $(shell go list ./... | grep -v /vendor/)

run:
	go run cmd/main.go

run_docker:
	docker run -it --rm -v /var/run/docker.sock:/var/run/docker.sock --group-add ${DOCKER_GROUP} -e "DOCKER_HOST:unix:///var/run/docker.sock" nicholasjackson/cnitch -hostname=docker

build_linux:
	CGO_ENABLED=0 GOOS=linux go build -o ./cmd/cnitch ./cmd/main.go

build_docker: build_linux
	docker build -t nicholasjackson/cnitch .
