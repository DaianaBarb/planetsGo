FROM golang:1.13

#COPY . /projeto-star-wars-api-go/

#COPY go.mod go.sum ./
#RUN go mod download


#RUN go mod download
#
#
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o star-wars-api-go ./cmd/main.go
RUN mkdir -p /opt/app

COPY api /opt/app/api

EXPOSE 8080
WORKDIR /opt/app

CMD ["./api"]
