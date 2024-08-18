build:
	go build -o exporter-boilerplate
fmt:
	go fmt github.com/coolapso/prometheush-exporter-boilerplate/...
	
build-docker-multiarch:
	docker build --platform linux/arm/v7,linux/arm64/v8,linux/amd64 -t exporter-boilerplate .
