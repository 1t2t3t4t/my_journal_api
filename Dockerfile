FROM golang:1.18beta1-buster AS build

WORKDIR /app
COPY ./go.mod .
COPY ./go.sum .

RUN go mod download

COPY . .

RUN go build .

ENTRYPOINT [ "./my_journal_api" ]
