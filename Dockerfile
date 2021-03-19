# Start from the latest golang base image
FROM golang:latest

RUN mkdir -p /go/src/app

WORKDIR /go/src/app

COPY . /go/src/app
RUN go mod download

RUN go build -o find-course ./cmd/app

EXPOSE 3000

CMD ["/go/src/app/find-course"]