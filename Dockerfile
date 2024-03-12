FROM golang:1.21 AS builder

LABEL maintainer = "Giwa Oluwatobi, giwaoluwatobi@gmail.com"

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -tags netgo -ldflags "-s -w" -o /app/Panth3r cmd/Panth3r/main.go

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=builder /app/Panth3r /Panth3r

ENTRYPOINT [ "/Panth3r" ]