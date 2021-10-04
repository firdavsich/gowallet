# prebuild
FROM golang:1.17-alpine3.14 AS builder

WORKDIR /go/src/app

# manage dependencies
COPY go.mod go.sum ./
RUN go mod download 
# Build
COPY . .
RUN go build 

# create image
FROM alpine:3.14

COPY --from=builder /go/src/app/gowallet /usr/local/bin/

CMD ["gowallet"]
