FROM golang:latest
LABEL maintainer="Denzyl Dick <hello@denzyl.io>"
#COPY go.mod go.sum ./
#RUN go mod download
COPY ./src ./src
RUN go get gopkg.in/mgo.v2
WORKDIR ./src
RUN go build -o api .
EXPOSE 80
CMD ["./api"]