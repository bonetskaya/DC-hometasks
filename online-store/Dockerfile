FROM golang:1.14.0-alpine3.11

COPY . /
WORKDIR /

RUN apk add git
RUN go get -u github.com/gorilla/mux     
RUN go build -o main .


# Run the executable
CMD ["./main"]

EXPOSE 8000
