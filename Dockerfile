FROM golang:1.18.2-alpine as builder

RUN apk add --no-cache git curl openssl

RUN mkdir -p /workspace/openpitrix-jobs/
WORKDIR /workspace/openpitrix-jobs/
COPY . .

RUN mkdir -p /release_bin
#RUN cd cmd/dump-all/ && CGO_ENABLED=0 GOBIN=/release_bin go install -mod=vendor -ldflags '-w -s'  d3os.io/openpitrix-jobs/cmd/dump-all/...
RUN CGO_ENABLED=0 GOBIN=/release_bin go install -mod=vendor -ldflags '-w -s'  d3os.io/openpitrix-jobs/cmd/import-app/...
#RUN CGO_ENABLED=0 GOBIN=/release_bin go install -mod=vendor -ldflags '-w -s'  d3os.io/openpitrix-jobs/cmd/upgrade/...

FROM alpine:3.16.0

WORKDIR /root
COPY import-config.yaml d3os/
COPY --from=builder /release_bin/* /usr/local/bin/

# Disable cache, always download chart package
ARG BUILDDATE
RUN echo "$BUILDDATE"
COPY urls.txt /root
RUN mkdir -p package && cp urls.txt package  && cd /root/package && for pkg in $(cat urls.txt); do wget $pkg; done && rm urls.txt
