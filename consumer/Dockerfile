# base image
FROM golang:1.19.3-alpine as builder

# create appuser.
RUN adduser -D -g '' appuser

# create workspace
WORKDIR /build/api

# download dependencies
COPY go.mod go.sum ./
RUN go mod download && \
  go mod verify

# copy source code
COPY . .

# build binary
RUN CGO_ENABLED=0 GOOS=linux \ 
    go build -o /go/bin/consumer-restapp ./cmd/restapp

# build a small image
FROM alpine

LABEL language="golang"

#WORKDIR /root

# use a non-root user
USER appuser

# import the user and group files from the builder
COPY --from=builder /etc/passwd /etc/passwd

# copy the static executable
COPY --from=builder /go/bin/consumer-restapp .

# expose the port
EXPOSE 3000

# run the app
CMD ["./consumer-restapp"]