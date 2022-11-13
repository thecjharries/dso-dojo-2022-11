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
ENV GIN_MODE=release
# This is because LocalStack only exposes port 22 on EC2 instances
# https://github.com/localstack/localstack/issues/6546
EXPOSE 22
ENV PORT=22
ENTRYPOINT ["./server"]
