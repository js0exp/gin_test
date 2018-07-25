FROM golang:1.9-alpine
WORKDIR /app
COPY ./gin_test /app/gin_test
EXPOSE 8000
CMD ["./gin_test"]