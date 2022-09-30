FROM golang:1.18-alpine3.15 AS builder

RUN apk --update add git make openssh

RUN addgroup -g 1001 -S builder && \
    adduser -u 1001 -S builder -G builder
USER builder:builder

ENV APP_HOME /home/builder
WORKDIR $APP_HOME

ARG SSH_PRIVATE_KEY
RUN mkdir -p ~/.ssh && echo "${SSH_PRIVATE_KEY}" > ~/.ssh/id_rsa && chmod 0600 ~/.ssh/id_rsa \
    && git config --global url."git@github.com:".insteadOf https://github.com/ \
    && ssh-keyscan github.com >> ~/.ssh/known_hosts

COPY --chown=builder:builder . $APP_HOME/

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# build executable to ./bin/server
RUN make build-native

FROM alpine:3.15

RUN addgroup -g 1001 -S runner && \
    adduser -u 1001 -S runner -G runner

ENV BUILD_DIR /home/builder
ENV APP_HOME /home/runner/bin
# create dummy .env file
RUN mkdir -p $APP_HOME && touch $APP_HOME/.env

USER runner:runner
WORKDIR $APP_HOME

COPY --chown=runner:runner --from=builder $BUILD_DIR/bin/ .
COPY --chown=runner:runner --from=builder go/bin/migrate .
COPY --chown=runner:runner --from=builder $BUILD_DIR/db/migrations ./migrations

ENTRYPOINT ["/home/runner/bin/gogen-project", "http-server"]
CMD ["--help"]
