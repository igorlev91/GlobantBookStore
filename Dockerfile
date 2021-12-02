FROM golang
COPY go.mod go.sum /go/src/github.com/igorlev91/GlobantBookStore/
WORKDIR /go/src/github.com/igorlev91/GlobantBookStore/
RUN go mod download
COPY *.go ./

RUN ls

RUN go get -u github.com/joho/godotenv
RUN go get -u gorm.io/gorm

RUN go get -u github.com/go-sql-driver/mysql
RUN go get gorm.io/driver/mysql
RUN go get -u github.com/gorilla/mux

EXPOSE 8080 8080
RUN go get github.com/githubnemo/CompileDaemon
RUN go mod download
ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main
