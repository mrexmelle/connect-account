FROM golang:1.19

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY cmd ./cmd
COPY internal ./internal
COPY config ./config
RUN go build -o ./connect-idp ./cmd/main.go

ENV APP_PROFILE docked
EXPOSE 8080
CMD ["/app/connect-idp serve"]