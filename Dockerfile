# Dockerfile
FROM golang:1.16

RUN mkdir app

RUN cd app

WORKDIR /app

COPY . .

ADD script app/

RUN go build

EXPOSE 8080
# Expose port 80
EXPOSE 80

CMD ["./main"]

