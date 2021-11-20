# go_server_image
FROM golang:1.17-buster
ENV PATH="${PATH}:$GOPATH"
WORKDIR /go/src
COPY . .
RUN go get github.com/githubnemo/CompileDaemon
CMD go mod download
ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main