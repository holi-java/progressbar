run:
	-@go run .

test:
	-@go test -v ./...

build: dist 
	-@go build -o dist .

clean:
	-@rm -rf dist

dist:
	-@mkdir dist
