FROM golang:1.24 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o client ./cmd/client/main.go

FROM gcr.io/distroless/base-debian11
COPY --from=builder /app/client /client
CMD ["/client"]