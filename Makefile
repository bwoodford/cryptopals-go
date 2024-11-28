build:
	go build -o bin/cryptopals-go main.go operations.go
clean:
	rm bin/cryptopals-go
