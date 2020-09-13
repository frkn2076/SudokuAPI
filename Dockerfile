FROM golang:latest

ENV GOPATH /go

FROM golang:latest
COPY ./Controllers /go/src/app/SudokuAPI/Controllers
COPY ./Helper /go/src/app/SudokuAPI/Helper
COPY ./Router /go/src/app/SudokuAPI/Router
COPY ./main.go /go/src/app/SudokuAPI/main.go

WORKDIR /go/src/app/SudokuAPI

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["SudokuAPI"]

