FROM golang:1.15
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main
CMD ["./main"]