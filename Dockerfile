
##
## Build
##

# go_server_image
FROM golang
ENV PATH="${PATH}:$GOPATH"
WORKDIR /go/src/source
COPY ./source .
RUN go get github.com/githubnemo/CompileDaemon
CMD go mod download
ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main



##
## Deploy
##
EXPOSE 8080


WORKDIR /

USER nonroot:nonroot

ENV BOOKSTORE_DBDRIVER="mysql"
ENV BOOKSTORE_DATABASE="root:@tcp(localhost:3306)/bookstore"
ENV BOOKSTORE_USER ='root'
ENV BOOKSTORE_PASSWORD =''
