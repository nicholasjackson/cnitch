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
	docker build -t quay.io/nicholasjackson/cnitch:latest .

push_docker:
	docker push quay.io/nicholasjackson/cnitch:latest

sonar_qube:
	sonar-scanner \
    -Dsonar.projectKey=cnitch \
    -Dsonar.sources=. \
    -Dsonar.host.url=https://sonarcloud.io \
    -Dsonar.organization=nicholasjackson-github \
    -Dsonar.login=abcdef0123456789
	19a5ee8d5578f46b534b4f8c112b25700285274a
