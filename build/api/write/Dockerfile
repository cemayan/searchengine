FROM golang:1.22.4-alpine3.20 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY configs .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o /app/writeapi  cmd/api/write/main.go

FROM busybox
COPY --from=builder /app/writeapi /app/writeapi
COPY --from=builder /app/configs /configs
EXPOSE 8088
ENTRYPOINT ["/app/writeapi"]