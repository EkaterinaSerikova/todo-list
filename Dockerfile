FROM golang:1.24-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o todolist cmd/todolist/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/todolist .
CMD ["./todolist"]
EXPOSE 8080