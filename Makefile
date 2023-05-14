
PROJECT_NAME=connect-idp
VERSION=0.2.0
DOCKER_BUILD=docker build
IMAGE_NAME=ghcr.io/mrexmelle/connect-idp

$(PROJECT_NAME): 
	go build -o $(PROJECT_NAME) cmd/*.go

clean:
	rm -f $(PROJECT_NAME)

distclean:
	rm -rf $(PROJECT_NAME) data

docker-image:
	docker build -t $(IMAGE_NAME):$(VERSION) .

docker-push:
	docker push $(IMAGE_NAME):$(VERSION)

test:
	go test ./internal/...

all: $(PROJECT_NAME)