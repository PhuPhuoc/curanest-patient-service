FROM golang:1.23.2 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
