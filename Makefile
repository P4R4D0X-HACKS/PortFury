build:
    go build -o bin/portfury ./cmd/portfury

run:
    go run ./cmd/portfury -host=scanme.nmap.org -ports=1-1000

test:
    go test ./...

clean:
    rm -rf bin/