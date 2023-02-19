# Build Stage
# First pull Golang image
FROM golang:1.20

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build ./cmd -v -o /url-shortener-backend

CMD ["/url-shortener-backend"]
