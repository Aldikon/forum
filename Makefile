build:
	docker build -t forum .

run-img:
	docker run --name forum -p 8080:8080 --rm forum

run:
	go run cmd/main.go