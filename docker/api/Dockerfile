FROM golang:1.16

ENV GO111MODULE=on

WORKDIR /go/src/github.com/satorunooshie/user-rest-api

COPY go.mod go.sum ./
# cache
RUN go mod download

COPY . .
RUN go build .

# hot reload
RUN go get github.com/pilu/fresh

EXPOSE 8080

CMD ["fresh"]
