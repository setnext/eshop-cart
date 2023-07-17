FROM golang:1.19-alpine

WORKDIR /app

COPY . .
COPY go.mod ./
COPY go.sum ./
COPY .env ./

RUN go mod download

COPY *.go ./

RUN go build -o /eshop-cart-api

EXPOSE 5000

CMD [ "/eshop-cart-api" ]
