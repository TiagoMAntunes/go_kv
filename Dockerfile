FROM golang:1.18.6
WORKDIR app
COPY . .
CMD ["go", "run", "server.go"]