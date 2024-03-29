FROM golang:alpine as builder

COPY . /go/src//tmp/tsstatus
WORKDIR /go/src//tmp/tsstatus

RUN set -ex \
 && apk add --update git \
 && go install \
      -ldflags "-X main.version=$(git describe --tags --always || echo dev)" \
      -mod=readonly

FROM alpine:latest

LABEL maintainer "Knut Ahlers <knut@ahlers.me>"

RUN set -ex \
 && apk --no-cache add \
      ca-certificates

COPY --from=builder /go/bin/tsstatus /usr/local/bin/tsstatus

EXPOSE 3000

ENTRYPOINT ["/usr/local/bin/tsstatus"]
CMD ["--"]

# vim: set ft=Dockerfile:
