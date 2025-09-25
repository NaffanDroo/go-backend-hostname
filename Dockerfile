##
## Build
##
FROM golang:1.20-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go  ./
COPY api api

RUN go build -o /hostname_api


##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /
COPY --from=build /hostname_api /hostname_api
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/hostname_api"]