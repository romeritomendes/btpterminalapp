// Copyright (c) 2026 Romerito Mendes Silva
// Licensed under the GPLv3. See LICENSE for details.
FROM docker.io/library/golang:alpine3.22 AS builder

WORKDIR /app

COPY server/go.mod server/go.sum ./

RUN go mod download

COPY server/ .

RUN CGO_ENABLED=0 GOOS=linux go build -o /btpterminalapp ./cmd/app

FROM debian:bookworm-slim

RUN useradd -r -u 10001 -m -d /home/appuser -s /usr/sbin/nologin appuser

RUN mkdir -p /app/.ssh \
 && chown -R appuser:appuser /app/.ssh \
 && chmod 700 /app/.ssh

WORKDIR /app
COPY --from=builder /btpterminalapp /btpterminalapp

USER appuser
EXPOSE 8080
CMD ["/btpterminalapp"]
