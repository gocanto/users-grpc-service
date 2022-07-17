############################
# STEP 1 build executable binary
############################

FROM golang:alpine as builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk add --no-cache git

WORKDIR $GOPATH/src/github.com/gocanto/gRPC/
COPY . .

# Fetch dependencies
RUN go mod tidy

# Build the binary. for grpc gateway
RUN go build ./server/server

RUN pwd
RUN echo $GOPATH

#EXPOSE 9090
# Run the hello binary.
#ENTRYPOINT ["./cmd/server"]

# final build
FROM alpine:3.16.0

WORKDIR /root/

COPY --from=builder /go/src/github.com/gocanto/rpc/server .

ENTRYPOINT ["bash", "-c", "/root/server"]
