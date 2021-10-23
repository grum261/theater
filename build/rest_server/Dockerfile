FROM golang:1.17.2-alpine3.14 AS builder

WORKDIR /build/

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags "-s -w" \ 
    github.com/grum261/theater/cmd/rest_server

RUN CGO_ENABLED=0 GOOS=linux go install -a -installsuffix cgo -ldflags "-s -w" -tags 'pgx' \
    github.com/golang-migrate/migrate/v4/cmd/migrate@v4.14.0

#-

FROM alpine:3.14 AS certificates

RUN apk --no-cache add ca-certificates

#-

FROM scratch

WORKDIR /api/

ENV PATH=/api/bin/:$PATH

COPY --from=certificates /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=builder /build/rest_server ./bin/rest_server
COPY --from=builder /go/bin/migrate ./bin/migrate
COPY --from=builder /build/internal/migrations .
COPY --from=builder /build/.env.dev .

EXPOSE 8000

CMD ["rest_server", "-env", "/api/.env.dev"]