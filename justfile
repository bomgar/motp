build:
    mkdir -p build
    go build -o build/motp main.go

clean:
    rm -rf build

check:
    go vet ./...
    golangci-lint run

fmt:
    go fmt ./...

test:
    go test -v ./...

