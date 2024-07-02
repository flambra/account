FROM golang:1.22-alpine AS builder

RUN apk add --no-cache tzdata

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd

FROM scratch

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /app /app
COPY --from=builder /app/.env /.env

ENV TZ=America/Sao_Paulo

EXPOSE 8081

CMD ["/app/main"]