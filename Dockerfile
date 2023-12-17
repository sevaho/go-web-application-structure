# vim: set ft=dockerfile:

FROM golang:alpine as builder

MAINTAINER Sebastiaan Van Hoecke

# Ca-certificates is required to call HTTPS endpoints in scratch images.
# Tini-static, https://github.com/krallin/tini init process, it ensures that the default signal handlers work
RUN apk --update add ca-certificates tini-static git

# Create appuser
ENV USER=appuser
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

# Golang needs a to be ran in a valid GOPATH
WORKDIR $GOPATH/src/app/
COPY . .

# Fetch dependencies.
RUN go get -d -v

# Build the binary for static target
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-extldflags "-static"' -a \
    -o /main .

# EXPERIMENTAL, makes the binary really small, but there is cost in performance.
# See https://upx.github.io/
# RUN apk --update add upx
# RUN upx /main 

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /main /main
COPY --from=builder /sbin/tini-static /sbin/tini-static
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Use an unprivileged user.
USER appuser:appuser

ENTRYPOINT ["/sbin/tini-static", "--", "/main"]
CMD ["--serve"]
