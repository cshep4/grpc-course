FROM golang:1.22.1-alpine as build

WORKDIR /build

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -gcflags='-N -l' -ldflags "-s -w" -o /hello-server ./cmd/server

FROM alpine as release

COPY --from=build /hello-server /hello-server

EXPOSE 50051
ENTRYPOINT ["/hello-server"]
