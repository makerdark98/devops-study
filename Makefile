main:
	CGO_ENABLED=0 go build -o main main.go

docker: main
	docker build -t makerdark98/helloworld:0.0.1 .
