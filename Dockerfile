FROM golang:alpine AS builder
ENV CGO_ENABLED=1
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN apk add --no-cache build-base
RUN GO111MODULE="on" GOOS=linux GOARCH=amd64 go build -o app ./cmd

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
COPY --from=builder /app/static static
RUN mkdir db

EXPOSE 8080
CMD ["./app"]