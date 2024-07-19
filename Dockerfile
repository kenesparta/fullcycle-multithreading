FROM golang:1.22-bookworm as builder
WORKDIR /app
RUN apt-get update && apt-get install -y ca-certificates
COPY . .
RUN make init && make build

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
