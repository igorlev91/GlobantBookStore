
RUNPATH=source/

BINARY_NAME=bookstore
TARGET_OS=macos
OUTPUT_DIR=${GOPATH}/bin


# Builds bookstore
build:
	GOARCH=amd64 GOOS=${TARGET_OS} go build -o '${OUTPUT_DIR}/${BINARY_NAME}-${TARGET_OS}' GlobantBookStore/source

# Runs the database container
run_db:
	docker run --name some-mysql -p 3306:3306 -e MYSQL_ALLOW_EMPTY_PASSWORD=true -e MYSQL_DATABASE=bookstore -d mysql:latest

# Runs bookstore
run: run_db
	GOARCH=amd64 GOOS=${TARGET_OS} go run GlobantBookStore /port 8080