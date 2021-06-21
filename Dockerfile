FROM golang:1.12-alpine AS build_base
RUN apk add --no-cache git
WORKDIR /tmp/scanner

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go test ./...
RUN go build -o ./out/dockerfile-scanner .

FROM alpine:3.9
RUN apk add ca-certificates

COPY --from=build_base /tmp/scanner/out/dockerfile-scanner /app/scanner
