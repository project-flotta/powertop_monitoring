FROM golang:1.18 as builder

USER root

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN apt-get update -y && \
    apt-get install -y curl && \
    # apt-get install -y libbtrfs-dev libgpgme-dev libdevmapper-dev &&\
    apt-get install -y powertop

RUN go build -o main ./cmd/
EXPOSE 8887
ENTRYPOINT ["/app/main"]
