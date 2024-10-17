FROM --platform=$BUILDPLATFORM minixxie/golang:1.21.0 as golang

ARG BUILDPLATFORM

WORKDIR /usr/local/go/src/app

COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

ADD . .
RUN ./scripts/build.sh
RUN cp ./bin/app /tmp/

FROM gcr.io/distroless/base
COPY --from=golang /tmp/app /app

ADD ./nodejs-project /tmp/nodejs-project

ENTRYPOINT ["/app"]
