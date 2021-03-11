FROM golang:1.16
WORKDIR /pokeapi
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
EXPOSE 3000
COPY . ./
RUN go build main.go
CMD ["./main"]