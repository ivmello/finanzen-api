ARG SERVICE_PATH

# STAGE 1: Prepare and Build
#
FROM golang:1.20-alpine AS base
WORKDIR /app
ADD . /app
RUN go mod vendor \
    && go mod download

#
# STAGE 2: Release
#
FROM base AS development
ARG SERVICE_PATH
ARG DEBUG_PORT
RUN apk add build-base \
    && go install github.com/go-delve/delve/cmd/dlv@latest \
    && go install github.com/cosmtrek/air@latest \
    && sed "s/{SERVICE_PATH}/${SERVICE_PATH}/g; s/{DEBUG_PORT}/${DEBUG_PORT}/g" ./docker/go-webserver/air.toml > /root/.air.toml
ENTRYPOINT ["air", "-c", "/root/.air.toml"]