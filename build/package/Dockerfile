FROM golang:1.19.3-alpine3.16

RUN mkdir /app

WORKDIR /app
COPY . /app

RUN go build -o archor .

CMD ["/app/archor"]
