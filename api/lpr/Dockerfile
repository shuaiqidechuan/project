FROM golang:alpine
WORKDIR /api
COPY . .
RUN go build -o lprapi ./lpr
CMD ["/api/lprapi"]