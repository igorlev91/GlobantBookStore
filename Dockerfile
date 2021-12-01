FROM golang
COPY go.mod go.sum /go/src/github.com/igorlev91/GlobantBookStore/
WORKDIR /go/src/github.com/igorlev91/GlobantBookStore/


RUN go get github.com/githubnemo/CompileDaemon
RUN go mod download
ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main

