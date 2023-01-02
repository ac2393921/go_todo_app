FROM golang:1.19.4-alpine3.17

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflogs "-w -s" -o app

# デプロイ用
FROM debian:bullseye-slim as deploy

RUN apt-get update

COPY --from=deploy-builder /app/app .

CMD ["./app"]

FROM golang:1.19.4-alpine3.17 as dev

# ローカル環境で利用するホットリロード環境
WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

RUN ls -la

CMD ["air", "-c", ".air.toml"]

