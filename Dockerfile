FROM 	golang:1.17.10-alpine3.15 as builder
RUN	mkdir /build && \
	cd /build
WORKDIR "/build"
COPY	. ./
RUN	CGO_ENABLED=0 go build -ldflags="-s -w" -o github-utility

FROM	alpine:3.15 
COPY	--from=builder /build/github-utility .
RUN	addgroup -S appgroup && adduser -S appuser -G appgroup
USER	appuser
ENTRYPOINT [ "./github-utility" ]
