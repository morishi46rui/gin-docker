# GolangのAlpineバージョンを使用
FROM golang:1.23-alpine3.21

# 環境変数
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOPROXY=https://proxy.golang.org,direct

# 作業ディレクトリの設定
WORKDIR /app

# 依存関係の取得
COPY go.mod go.sum ./
RUN go mod download

# 🔥 `docs/` を明示的にコピー
COPY docs ./docs
COPY main.go ./

# ビルド
RUN GOOS=linux GOARCH=amd64 go build -o gin-server main.go && ls -l gin-server

# ポートの公開
EXPOSE 8080

# コンテナ起動時に実行するコマンド
CMD ["/app/gin-server"]
