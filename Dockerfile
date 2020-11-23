FROM golang:latest

WORKDIR /app
COPY . /app/gin-blog

RUN go build .

EXPOSE 8000

ENTRYPOINT ["./gin-blog"]