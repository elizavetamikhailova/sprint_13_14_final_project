FROM golang:1.23

WORKDIR /app

RUN apt-get -y update
RUN apt-get -y upgrade

RUN apt-get install -y sqlite3 libsqlite3-dev

COPY go.mod go.sum ./

RUN go mod download

COPY ./ ./

RUN /usr/bin/sqlite3 scheduler.db

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /my_app

CMD ["/my_app"]