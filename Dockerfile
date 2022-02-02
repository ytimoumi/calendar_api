FROM golang:1.16-alpine As builder
RUN apk update && apk add --no-cache git ca-certificates tzdata bash mysql-client


# Set the time zone
RUN cp /usr/share/zoneinfo/Europe/Paris /etc/localtime
RUN echo "Europe/Paris" >  /etc/timezone

WORKDIR /go/src/giskard/

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go get -v github.com/pilu/fresh


ENTRYPOINT fresh