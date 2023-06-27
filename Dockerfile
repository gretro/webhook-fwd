FROM golang:1.20.5-alpine AS deps

WORKDIR /webhook-fwd

COPY go.mod go.sum ./

RUN apk add --no-cache make ca-certificates bash \
  && update-ca-certificates \
  && go mod download

FROM deps AS code

WORKDIR /webhook-fwd

COPY ./ ./

FROM code AS builder

WORKDIR /webhook-fwd

RUN make build_api

FROM alpine:3.18 AS runner

WORKDIR /webhook-fwd

COPY --from=builder --chown=api:api --chmod=rx /webhook-fwd/dist/ ./

USER api

EXPOSE 5333

ENTRYPOINT [ "/webhook-fwd/webhook-fwd-api" ]
