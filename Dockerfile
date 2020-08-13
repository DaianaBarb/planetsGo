FROM golang:1.14
WORKDIR /projeto-star-wars-api-go/
COPY ./ .
RUN  go build /projeto-star-wars-api-go/cmd/projeto-star-wars-api-go
CMD ["/projeto-star-wars-api-go/projeto-star-wars-api-go"]  
