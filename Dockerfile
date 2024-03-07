FROM golang:1.22-alpine3.19 AS builder
RUN apk add git make

WORKDIR /src
# download and cache go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN make build

FROM alpine:3.19
RUN apk add --no-cache ca-certificates
RUN update-ca-certificates
COPY --from=builder /src/foosvc /usr/local/bin/app

EXPOSE 80
CMD ["/usr/local/bin/app", "server"]