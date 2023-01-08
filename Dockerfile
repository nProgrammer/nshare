FROM golang:1.19.2

WORKDIR /api
COPY go.mod .
RUN go mod tidy
RUN go mod download

COPY main.go .

RUN go build -o /nshare-api

WORKDIR /
RUN rm -r /api

CMD [ "/nshare-api" ]