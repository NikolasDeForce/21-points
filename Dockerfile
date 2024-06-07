FROM golang:latest

COPY . /app

EXPOSE 8010
WORKDIR /app/server
ENTRYPOINT [ "go", "run", "server.go" ]