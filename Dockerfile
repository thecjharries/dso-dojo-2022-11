FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN apk add --no-cache git make
RUN make build


FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/bin/server .
EXPOSE 8080
ENV GIN_MODE=release
ENTRYPOINT ["./server"]
