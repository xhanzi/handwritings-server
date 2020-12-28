FROM golang:alpine AS builder
WORKDIR /opt
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-w'

FROM alpine:edge
WORKDIR /opt/
COPY --from=builder /opt/handwritings-server .
EXPOSE 8090
ENTRYPOINT [ "/opt/handwritings-server" ]
