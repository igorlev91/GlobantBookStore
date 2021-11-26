##
## Build
##
FROM go1.17.3-buster


ENV PATH="${PATH}:$GOPATH"
WORKDIR /go/src/GlobantBookStore

COPY . .
RUN go get github.com/githubnemo/CompileDaemon
CMD go mod download
ENTRYPOINT CompileDaemon --build="go build runserver.go" --command=./runserver


##
## Deploy
##

ENV BOOKSTORE_DBDRIVER="mysql"
ENV BOOKSTORE_DATABASE="root:@tcp(localhost:3306)/bookstore"

