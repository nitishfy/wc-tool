FROM golang:1.21-alpine as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o wc .
FROM alpine:3.11.3

COPY --from=build /app/wc /app/wc
WORKDIR /app
CMD ["./wc"]
