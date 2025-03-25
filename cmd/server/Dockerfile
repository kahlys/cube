FROM golang:1.24 AS builder
WORKDIR /app
COPY go.mod go.sum ./
COPY . .
RUN go build -o server ./cmd/server/main.go

FROM gcr.io/distroless/base-debian11
COPY --from=builder /app/server /server
EXPOSE 50051
CMD ["/server"]