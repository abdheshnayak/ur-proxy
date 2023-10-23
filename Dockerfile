# syntax=docker/dockerfile:1.4
FROM golang:1.18.3-alpine3.16 AS base
USER 1001
ENV GOPATH=/tmp/go
ENV GOCACHE=/tmp/go-cache
WORKDIR /tmp/app
COPY --chown=1001 --from=project-root ./go.mod ./go.sum ./tools.go ./
RUN go mod download -x
ARG APP
RUN mkdir -p ./$APP
WORKDIR /tmp/app/$APP
COPY --chown=1001 ./  ./
RUN CGO_ENABLED=0 go build -tags musl -o /tmp/bin/$APP ./main.go
RUN chmod +x /tmp/bin/$APP

FROM gcr.io/distroless/static-debian11
ARG APP
COPY --from=base --chown=1001 /tmp/bin/$APP /ur-proxy
COPY templates /templates
CMD ["/ur-proxy"]

