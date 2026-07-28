FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY internal/ internal/
COPY main.go ./

RUN go build -o /app/optimizer-server main.go

FROM alpine:3.19

WORKDIR /app
COPY --from=builder /app/optimizer-server .

EXPOSE 8080

CMD ["./optimizer-server"]
