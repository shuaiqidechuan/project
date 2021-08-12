FROM golang:alpine
WORKDIR /api
COPY . .
RUN go build -o shiborapi ./shibor
CMD ["/api/shiborapi"]