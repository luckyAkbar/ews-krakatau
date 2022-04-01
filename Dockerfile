FROM golang:1.18-alpine

WORKDIR /app

RUN mkdir src
RUN mkdir bin
RUN mkdir data

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go mod tidy

WORKDIR /app/src

COPY . .

RUN go build -o /app/bin/main main.go

EXPOSE 3333

WORKDIR /app

RUN rm -r src/

CMD ["./bin/main", "server"]