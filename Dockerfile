# Use an official Go runtime as a parent image
FROM golang:1.21rc2-alpine3.18

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# prefetch the binaries, so that they will be cached and not downloaded on each change
# https://goprisma.org/docs/reference/deploy/docker
RUN go run github.com/steebchen/prisma-client-go prefetch

COPY . /app

RUN go get -d -v ./...

RUN go run github.com/steebchen/prisma-client-go generate

RUN go build -o main ./api

EXPOSE 8080

CMD ["./main"]
