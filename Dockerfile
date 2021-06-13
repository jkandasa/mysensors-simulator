FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.16-alpine3.13 AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app

RUN apk add --no-cache git

ARG GOPROXY
# download deps before gobuild
RUN go mod download -x

ARG TARGETOS
ARG TARGETARCH
RUN source ./scripts/version.sh && \
  GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH:-amd64} go build -v -o ms-simulator -ldflags "$LD_FLAGS" cmd/main.go

FROM alpine:3.13

LABEL maintainer="Jeeva Kandasamy <jkandasa@gmail.com>"

ENV APP_HOME="/app"

# install timzone utils
RUN apk --no-cache add tzdata

# create a user and give permission for the locations
RUN mkdir -p ${APP_HOME}

# copy application bin file
COPY --from=builder /app/ms-simulator ${APP_HOME}/ms-simulator

RUN chmod +x ${APP_HOME}/ms-simulator

# copy default files
COPY ./resources/config.yaml ${APP_HOME}/config.yaml

WORKDIR ${APP_HOME}

CMD ["/app/ms-simulator", "-config", "/app/config.yaml"]
