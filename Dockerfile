FROM golang:alpine
WORKDIR /project
COPY . .
RUN go build -o exec ./api/cmd
CMD ["/project/exec"]