FROM golang:1.14-alpine AS builder
ADD . /workspace/
WORKDIR /workspace/
# TODO: ldflags should get container environmental variables or args passed by callers.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "sampleApp" -ldflags="-X 'main.version=v1.0.0' -X 'main.date=$(date "+%Y-%m-%dT%H:%M:%SZ")'"

FROM alpine:latest
WORKDIR /root
RUN apk --no-cache --update add ca-certificates tzdata && update-ca-certificates
COPY --from=builder /workspace/  /usr/local/bin/
ENTRYPOINT ["sampleApp"]
