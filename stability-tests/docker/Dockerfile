ARG SEDRAD_IMAGE
ARG SEDRAMINER_IMAGE

FROM ${SEDRAD_IMAGE} as serad
FROM ${SEDRAMINER_IMAGE} as seraminer

FROM golang:1.19-alpine

RUN mkdir -p /go/src/github.com/seracoin/serad

WORKDIR /go/src/github.com/seracoin/serad

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=serad /app/ /app/
COPY --from=seraminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/seracoin/serad/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]