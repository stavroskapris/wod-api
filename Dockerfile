# golang image
FROM golang:1.14-alpine 

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 8081
