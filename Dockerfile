# Dockerfile
FROM golang:1.16

WORKDIR /app

COPY . .

RUN go build

EXPOSE 8080
# Expose port 80
EXPOSE 80

CMD ["./main"]

