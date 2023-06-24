FROM golang:alpine as builder

RUN apk add git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd/purl
RUN go build -o /purl

FROM alpine:latest as release


WORKDIR /app/

COPY --from=builder /purl .

EXPOSE 1378

ENTRYPOINT ["./purl"]

