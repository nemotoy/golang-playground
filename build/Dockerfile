FROM golang:1.19.3-alpine AS builder
ARG VERSION
ADD . /workspace/
WORKDIR /workspace/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "sampleApp" -ldflags="-X 'main.version=$VERSION' -X 'main.date=$(date "+%Y-%m-%dT%H:%M:%SZ")'"

FROM alpine:latest
WORKDIR /root
RUN apk --no-cache --update add ca-certificates tzdata && update-ca-certificates
COPY --from=builder /workspace/  /usr/local/bin/
ENTRYPOINT ["sampleApp"]
