FROM golang:1.21-alpine AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o . .

FROM alpine:3.13
RUN apk add bash
RUN apk add curl
RUN apk --no-cache add ca-certificates tzdata \
    && cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime 
RUN mkdir -p /var/www/html/golang
WORKDIR /var/www/html/golang
COPY --from=builder /app/dating-api .
COPY config.yml /var/www/html/golang/config.yml
EXPOSE 8080
CMD ["./dating-api", "s"]