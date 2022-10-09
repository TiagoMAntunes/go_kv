FROM golang:1.18.6
WORKDIR app
# EXPOSE 3000j
COPY . .
CMD ["go", "run", "server.go"]