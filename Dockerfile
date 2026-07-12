# syntax=docker/dockerfile:1

FROM golang:1.26-alpine AS builder

RUN apk add --no-cache ca-certificates git

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /bin/server ./cmd/server

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /bin/server /server

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/server"]
