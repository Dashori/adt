FROM debian:12.1-slim

RUN apt update 
RUN apt show golang
RUN apt install golang ca-certificates -y

COPY . /

WORKDIR /backend

RUN go mod download

RUN go build src/cmd/main.go

EXPOSE 8080

CMD ["./main"]
