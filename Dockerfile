FROM golang:1.19.2

WORKDIR /api
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY main.go .
ADD certificates certificates
ADD errors errors
ADD repositories repositories
ADD models models
ADD responses responses

RUN go build -o /nshare-api

WORKDIR /
RUN rm -r /api

ENV WORKINGPORT=8080
ENV SSLMODE=false
ENV DB_SERVERNAME=localhost
ENV DB_SERVER_PORT=3309
ENV DB_NAME=nShare
ENV DB_PASSWD=example
ENV DB_USER=root

CMD [ "/nshare-api" ]