FROM golang:1.20.5-bookworm
# Note on image tags: golang:*-alpine* has no git, which go-get uses,
# so unless you `apk add git`, it might throw a related error.

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Prefetch the binaries, so that they will be cached and not downloaded on each change
# (see https://goprisma.org/docs/reference/deploy/docker):
RUN go run github.com/steebchen/prisma-client-go prefetch

COPY . /app

RUN go get -d -v ./...

RUN go run github.com/steebchen/prisma-client-go generate

RUN go test ./api
RUN go build -o main ./api

EXPOSE 8080

CMD ["./main"]
