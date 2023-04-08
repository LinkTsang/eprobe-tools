FROM golang:1.19

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY app ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /simple-http-server

EXPOSE 8080

CMD ["/simple-http-server"]