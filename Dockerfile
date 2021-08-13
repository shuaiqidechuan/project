FROM golang:alpine
WORKDIR /project
COPY . .
RUN go build -o exec ./api
CMD ["/project/exec"]