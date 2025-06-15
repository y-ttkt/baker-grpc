# protoc を実行するのは Go 製 API サービス（api/ ディレクトリ）をビルドするときだけ
FROM golang:1.21.5
ENV TZ=Asia/Tokyo

# protoc 本体
RUN apt-get update && \
    apt-get install -y --no-install-recommends protobuf-compiler && \
    rm -rf /var/lib/apt/lists/*

# Go 用プラグイン
# protoc-gen-go (v1.31.0) と protoc-gen-go-grpc (v1.5.1) を別RUNでインストール
RUN go install \
      google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0

RUN go install \
      google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
      protobuf-compiler \
      libprotobuf-dev && \
    rm -rf /var/lib/apt/lists/*

# PATH にバイナリを載せる
ENV PATH="$PATH:$(go env GOPATH)/bin"

WORKDIR /app

COPY api/go.mod  ./
COPY proto/ ./proto
RUN go mod download