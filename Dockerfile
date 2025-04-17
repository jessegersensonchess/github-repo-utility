FROM 	golang:1.23.8-alpine as builder
RUN	mkdir /build && \
	cd /build
WORKDIR "/build"
COPY	. ./
RUN	CGO_ENABLED=0 go build -ldflags="-s -w" -o github-utility

FROM	alpine:3
COPY	--from=builder /build/github-utility .
RUN	addgroup -S appgroup && adduser -S appuser -G appgroup
USER	appuser
ENTRYPOINT [ "./github-utility" ]
