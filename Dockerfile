FROM golang:1.22.0-alpine3.19 AS builder

ARG VERSION

WORKDIR /workspace

# add go modules lockfiles
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go install github.com/google/wire/cmd/wire@latest && wire ./internal/wire/wire.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o app .

# use the scratch image for the smallest possible image size
FROM alpine:3.19

ENV VERSION=${VERSION}

# copy over SSL certificates, so that we can make HTTPS requests
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /workspace/app /app

ENTRYPOINT ["/app"]
