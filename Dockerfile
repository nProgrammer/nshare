FROM golang:1.19.2

WORKDIR /api
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY main.go .
ADD certificates certificates
ADD errors errors
ADD repositories repositories

RUN go build -o /nshare-api

WORKDIR /
RUN rm -r /api

ENV WORKINGPORT=8080
ENV SSLMODE=false

CMD [ "/nshare-api" ]