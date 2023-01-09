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
ADD controllers controllers
ADD utils utils

RUN go build -o /nshare-api

WORKDIR /
RUN rm -r /api

CMD [ "/nshare-api" ]