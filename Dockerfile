FROM golang:1.19-alpine AS builder
RUN apk add --no-cache git make
WORKDIR /app
COPY ./go.mod .
COPY ./go.sum .
COPY ./main.go .
COPY ./Makefile .
COPY ./.git .
RUN go mod download
RUN make build


FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/bin/server .
EXPOSE 8080
ENV GIN_MODE=release
ENTRYPOINT ["./server"]
