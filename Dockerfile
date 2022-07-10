FROM golang:1.17.10-alpine3.15 as builder
RUN  	 mkdir /build && \
	 cd /build
WORKDIR "/build"
COPY . ./
RUN 	go build -o github-repo-utility

FROM alpine:3.15 
COPY --from=builder /build/github-repo-utility .
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser
ENTRYPOINT [ "./github-repo-utility" ]
