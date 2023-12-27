FROM golang:1-alpine AS builder
RUN apk add --no-cache build-base
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /dist/app ./cmd/ws-server
RUN ldd /dist/app | tr -s [:blank:] '\n' | grep ^/ | xargs -I % install -D % /dist/%


FROM scratch
COPY --from=builder /dist /
USER 65534
ENTRYPOINT ["/app"]