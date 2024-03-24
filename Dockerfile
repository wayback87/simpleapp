FROM golang:1.22 AS build

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/simpleapp

FROM alpine:3.19.1

WORKDIR /app

COPY --from=build /app/simpleapp ./simpleapp

ENTRYPOINT ["/app/simpleapp"]