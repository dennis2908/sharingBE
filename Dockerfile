FROM golang:1.16
WORKDIR /go/src/quickstart
COPY . . 
RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 9333
# Install server application
CMD ["go", "run", "main.go"]