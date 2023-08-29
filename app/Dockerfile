# FROM debian:12.1-slim

# RUN apt update 
# RUN apt show golang
# RUN apt install golang ca-certificates -y

# WORKDIR /

# COPY go.mod .
# COPY go.sum .

# RUN go mod download

# COPY . /

# RUN go build src/cmd/main.go

# EXPOSE 8080

# CMD ["./main"]

FROM debian:12.1-slim

RUN apt-get update && \
    apt-get install golang-go ca-certificates -y

COPY . /app

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go build src/cmd/main.go

EXPOSE 8080

CMD ["./main"]
