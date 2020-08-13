FROM golang:1.14
WORKDIR /projeto-star-wars-api-go/
COPY ./ .
RUN  go build /projeto-star-wars-api-go/cmd/main.go 
CMD ["/projeto-star-wars-api-go/cmd/main.go"]
