FROM --platform=$BUILDPLATFORM golang:latest AS builder
ARG TARGETARCH

WORKDIR /exporter
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} go build -a -o exporter-boilerplate

FROM alpine:latest

COPY --from=builder exporter/exporter-boilerplate /usr/bin/exporter-boilerplate

EXPOSE 9101
ENTRYPOINT ["/usr/bin/exporter-boilerplate"]
