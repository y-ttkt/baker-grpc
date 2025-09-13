FROM golang:1.23.5
ENV TZ=Asia/Tokyo

# protoc 本体
RUN apt-get update && \
    apt-get install -y --no-install-recommends protobuf-compiler && \
    rm -rf /var/lib/apt/lists/*

RUN go install \
      google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0

RUN go install \
      google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
      protobuf-compiler \
      libprotobuf-dev && \
    rm -rf /var/lib/apt/lists/*

ENV PATH="$PATH:$(go env GOPATH)/bin"

WORKDIR /app/api

COPY api/go.mod  ./
RUN go mod download