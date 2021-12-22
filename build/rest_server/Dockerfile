FROM golang:1.17.5-alpine3.15 AS builder

WORKDIR /build/

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags "-extldflags -static" -tags musl \ 
    -o ./rest ./cmd/rest_server

#-

FROM alpine:3.15 AS certificates

RUN apk --no-cache add ca-certificates

#-

FROM scratch

WORKDIR /api/

ENV PATH=/api/bin/:$PATH

COPY --from=certificates /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=builder /build/rest ./bin/rest
COPY --from=builder /build/.env .
COPY --from=builder /build/third_party .

EXPOSE 8000

CMD ["rest", "-env", "/api/.env", "-port", ":8000"]