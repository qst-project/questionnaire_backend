FROM golang:latest

WORKDIR /backend

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./backend

RUN go build -o /backend

EXPOSE 8080

CMD [ "/app" ]