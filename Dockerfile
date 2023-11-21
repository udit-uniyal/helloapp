# Dockerfile
FROM golang:1.16

WORKDIR /app

COPY . .

# Install cat command (assuming it's available in the package manager of the base image)
RUN apt-get update && \
    apt-get install -y cat

RUN go build

EXPOSE 8080

CMD ["./main"]
