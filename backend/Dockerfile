# Build Stage
# First pull Golang image
FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

WORKDIR /app/cmd
RUN go build -v -o /url-shortener-backend

CMD ["/url-shortener-backend"]
