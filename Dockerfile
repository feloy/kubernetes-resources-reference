FROM golang:1.13-buster as gobuilder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# Download dependencies
RUN go mod download

# Copy the go source
COPY main.go main.go
COPY api/ api/
COPY structure/ structure/
COPY markdown/ markdown/
COPY docbook/ docbook/

# Build
RUN GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o k8sref main.go

FROM eu.gcr.io/k8sref-io/docbook

WORKDIR /root

COPY --from=gobuilder /workspace/k8sref /usr/local/bin/
