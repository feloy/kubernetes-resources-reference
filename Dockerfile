FROM golang:1.13-buster AS builder

WORKDIR /app

ADD go.mod go.sum /app/
RUN go mod download

ADD api /app/api
ADD config /app/config

ADD pkg /app/pkg
ADD cmd /app/cmd

# Run unit tests
RUN go test ./...

# Build binary
RUN GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o /kubernetes-api-reference cmd/main.go

# Verify that resourceslist.txt is up to date
RUN /kubernetes-api-reference resourceslist -f api/v1.19/swagger.json > /tmp/resourceslist.txt
RUN diff /tmp/resourceslist.txt api/v1.19/resourceslist.txt

# final stage
FROM eu.gcr.io/k8sref-io/docbook

WORKDIR /root

COPY --from=builder /kubernetes-api-reference /usr/local/bin/
