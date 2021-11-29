
FROM golang
ENV PATH="${PATH}:$GOPATH"
WORKDIR /go/src/source
COPY ./source .
RUN go get github.com/githubnemo/CompileDaemon
CMD go mod download
ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main

