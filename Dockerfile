FROM golang:1.15
WORKDIR $GOPATH/src/connectfour
COPY . .
CMD ["go", "run", "main.go"]
