SRC = src
BIN = bin
main: test
	CGO_ENABLED=0 go build -o $(BIN)/main $(SRC)/main.go

docker: main
	docker build -t makerdark98/helloworld:0.0.1 .

test:
	CGO_ENABLED=0 go test $(SRC)/**
clean:
	rm -rf $(BIN)
