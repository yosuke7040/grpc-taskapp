# syntax=docker/dockerfile:1

FROM golang:1.22-bookworm AS builder

WORKDIR /workspace

COPY go.mod ./
# COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build

FROM golang:1.22-alpine3.19 AS runner

WORKDIR /workspace

COPY --from=builder /workspace ./

EXPOSE 8080

CMD ["sleep", "infinity"]
# CMD ["/workspace"]

FROM golang:1.22-bookworm AS develop

WORKDIR /workspace
# COPY --from=builder /workspace ./

# モック生成のためにvektra/mockeryをインストール
RUN curl -L https://github.com/vektra/mockery/releases/download/v2.42.1/mockery_2.42.1_Linux_x86_64.tar.gz | tar xvz && mv ./mockery /usr/bin/mockery

# データベースのマイグレーションのためにgolang-migrate/migrateをインストール
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz && mv ./migrate /usr/bin/migrate

# サーバーのライブリロードのためにcosmtrek/airをインストール
# RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/air-verse/air@v1.52.2

# データベースアクセスのためにkyleconroy/sqlcをインストール
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# gRPCとREST APIのコード生成のためにconnect-goとbufをインストール
RUN go install github.com/bufbuild/buf/cmd/buf@latest \
  && go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
  && go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest

# # モジュールへ実行権限を与える(vscodeユーザーでもアクセス可能にする)
# RUN sudo chmod -R a+rwX /go/pkg

# CMD ["air"]
CMD ["sleep", "infinity"]
