FROM golang:1.26.1-alpine

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /usr/local/bin/app ./cmd/app

CMD ["app"]