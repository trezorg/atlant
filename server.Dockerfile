FROM golang:alpine as builder
ENV USER=proxy APPNAME=atlant USER_ID=1000

RUN apk add make git protobuf-dev protobuf ca-certificates && \
    go get google.golang.org/protobuf/cmd/protoc-gen-go \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc && \
    adduser -D -H -u ${USER_ID} ${USER}

ADD go.mod /build/
RUN cd /build && go mod download

ARG VERSION=0.0.1

ADD . /build/
RUN cd /build && \
    export PATH="$PATH:$(go env GOPATH)/bin" && \
    VERSION=${VERSION} BINARY=${APPNAME} make build

FROM scratch
ENV USER=proxy APPNAME=atlant APPDIR=/app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/${APPNAME} ${APPDIR}/
COPY --from=builder /etc/passwd /etc/passwd
WORKDIR ${APPDIR}
USER ${USER}
ENTRYPOINT ["/app/atlant"]
