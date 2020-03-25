FROM golang:1.14.0-alpine3.11

RUN apk update && apk add --no-cache git

WORKDIR /

COPY . .

CMD ["go", "run", "main.go"]

EXPOSE 8000
