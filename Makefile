TARGET=tic-tac-toe
MAIN=main.go

build:
	@go build -o bin/${TARGET} ${MAIN}

run: build
	@./bin/${TARGET}

watch:
	@gow -c run ${MAIN}

test:
	@go test -v ./...
