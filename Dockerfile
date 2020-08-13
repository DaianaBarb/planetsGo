FROM golang:1.14

COPY . /projeto-star-wars-api-go/

WORKDIR /projeto-star-wars-api-go/

COPY go.mod go.sum ./
RUN go mod download


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o star-wars-api-go ./cmd/main.go

CMD ["./star-wars-api-go"]
