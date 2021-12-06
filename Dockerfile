FROM golang
ENV PATH="${PATH}:$GOPATH"
COPY go.mod go.sum /go/src/github.com/igorlev91/GlobantBookStore/


WORKDIR /go/src/github.com/igorlev91/GlobantBookStore/source

COPY ./source .


RUN go get github.com/githubnemo/CompileDaemon
RUN go mod download
ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main