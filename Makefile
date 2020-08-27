build:
	@go build -o bin/go-shortener-api -v
run:
	@export $(cat .env | xargs) && ./bin/go-shortener-api
clean:
	@go clean -o bin/go-shortener-api
vendor:
	@go mod vendor
test:
	@export $(cat .env | xargs) && go test -v