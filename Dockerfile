FROM golang:1.22-bookworm as builder
WORKDIR /app
COPY . .
RUN make init && make build

FROM scratch
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
