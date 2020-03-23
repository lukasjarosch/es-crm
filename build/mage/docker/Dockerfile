# syntax=docker/dockerfile:experimental
# Arguments
ARG BUILD_DATE
ARG VERSION
ARG VCS_REF

## BUILD IMAGE
FROM golang:1.13 as builder

# hadolint ignore=DL3008
RUN set -eux; \
    apt-get update -y && \
    apt-get install -y --no-install-recommends apt-utils bzr upx zip unzip;

# create a non-root privilege account to build
RUN adduser --disabled-password --gecos "" -u 1000 golang && \
    mkdir -p "$GOPATH/src/workspace" && \
    chown -R golang:golang "$GOPATH/src/workspace" && \
    mkdir /home/golang/.ssh && \
    mkdir /var/ssh && \
    chown -R golang:golang /home/golang && \
    chown -R golang:golang /var/ssh && \
    chmod 700 /home/golang

# force go modules
ENV GO111MODULE=on

# disable go proxy
ENV GOPROXY=direct
ENV GOSUMDB=off
WORKDIR $GOPATH/src/workspace

# prepare an unprivilegied user for run
RUN set -eux; \
    echo 'nobody:x:65534:65534:nobody:/:' > /tmp/passwd && \
    echo 'nobody:x:65534:' > /tmp/group

# switch to build user
USER golang

# copy build tools
COPY --chown=golang:golang mage.go .
COPY --chown=golang:golang tools tools

# clean go mod cache
RUN set -eux; \
	go clean -modcache

# install tools
#RUN set -eux; \
#    go run mage.go -d tools

# copy project go module
COPY --chown=golang:golang . .

# build final target
RUN set -eux; \
    go run mage.go -d services/{{.Service}}

# compress binaries
RUN set -eux; \
    upx -9 bin/* && \
    chmod +x bin/*

## -------------------------------------------------------------------------------------------------
# hadolint ignore=DL3007
FROM gcr.io/distroless/static:latest
# Arguments
ARG BUILD_DATE
ARG VERSION
ARG VCS_REF
# Metadata
LABEL \
    org.label-schema.build-date=$BUILD_DATE \
    org.label-schema.name="{{.Name}}" \
    org.label-schema.description="{{.Description}}" \
    org.label-schema.url="{{.URL}}" \
    org.label-schema.vcs-url="https://github.com/Zenithar/go-monorepo.git" \
    org.label-schema.vcs-ref=$VCS_REF \
    org.label-schema.vendor="Zenithar" \
    org.label-schema.version=$VERSION \
    org.label-schema.schema-version="1.0"
COPY --from=builder /go/src/workspace/bin/{{.Service}} /usr/bin/{{.Service}}
COPY --from=builder /tmp/group /tmp/passwd /etc/
USER nobody:nobody
WORKDIR /
ENTRYPOINT [ "/usr/bin/{{.Service}}" ]
CMD ["--help"]
