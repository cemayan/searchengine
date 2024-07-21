FROM golang:1.22.4-alpine3.20 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY configs .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o /app/readapi  cmd/api/read/main.go

FROM scratch
COPY --from=builder /app/readapi /app/readapi
COPY --from=builder /app/configs /configs
EXPOSE 8087
CMD ["/app/readapi"]