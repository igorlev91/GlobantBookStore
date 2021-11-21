# go_server_image
FROM golang:1.17-buster
ENV PATH="${PATH}:$GOPATH"
WORKDIR GlobantBookStore
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN ls
RUN pwd

RUN go mod download

RUN go get -u github.com/go-sql-driver/mysql
RUN go get gorm.io/driver/mysql
RUN go get -u github.com/gorilla/mux
Run go get -u https://github.com/igorlev91/GlobantBookStore

RUN go build -o GlobantBookStore
RUN go get github.com/githubnemo/CompileDaemon

CMD go mod download
ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main

