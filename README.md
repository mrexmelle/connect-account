# connect-idp

## Compiling

```
$ go build -o connect-idp cmd/*.go
```

## Building Docker image

Note that only the owner of the repository is allowed to build the image. 

```
$ docker build -t ghcr.io/mrexmelle/connect-idp:${VERSION} .
$ docker push ghcr.io/mrexmelle/connect-idp:${VERSION}
```

## Running

### For local environment

```
$ docker pull postgres:15-alpine
$ docker run \
	-v data:/var/lib/postgresql/data \
	-v init-db:/docker-entrypoint-initdb.d \
	-p 5432:5432 \
	-e POSTGRES_PASSWORD=123 \
	--restart always \
	postgres:15-alpine
$ ./connect-idp serve
```

### For docker environment

```
$ docker compose up
```
Note that you cannot alter the docker image in the container registry. Only the owner of the repository is allowed to do so.

If error happens in `core` service due to failure to connect to database, restart it:
```
$ docker compose restart core
```
The failure happens due to `db` service isn't ready when `core` attempts to connect to it.
