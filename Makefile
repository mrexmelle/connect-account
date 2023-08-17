
PROJECT_NAME=connect-idp
VERSION=0.2.1
IMAGE_NAME=ghcr.io/mrexmelle/$(PROJECT_NAME)

$(PROJECT_NAME): 
	go build -o $(PROJECT_NAME) cmd/*.go

clean:
	rm -f $(PROJECT_NAME)

distclean:
	rm -rf $(PROJECT_NAME) data

docker-image:
	docker build -t $(IMAGE_NAME):$(VERSION) .

docker-release:
	docker push $(IMAGE_NAME):$(VERSION)

test:
	go test ./internal/...

all: $(PROJECT_NAME)