# Vendor stage
FROM golang:1.23 as dep
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go mod vendor

# Build binary stage
FROM golang:1.23 as build
WORKDIR /build
COPY --from=dep /build .
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o server -tags nethttpomithttp2 ./cmd

# Minimal image
FROM alpine:latest
WORKDIR /app
COPY configs/config.yaml configs/config.yaml
COPY --from=build /build/server server
COPY migrations migrations
RUN apk update
RUN apk upgrade
RUN apk add ca-certificates
RUN apk --no-cache add tzdata
CMD ["./server"]
