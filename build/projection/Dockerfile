FROM golang:1.22.4-alpine3.20 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY configs .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o /app/projection  cmd/projection/main.go

FROM busybox
COPY --from=builder /app/projection /app/projection
COPY --from=builder /app/configs /configs
ENTRYPOINT ["/app/projection"]