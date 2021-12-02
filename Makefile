# Local
install:
	@go get -u ./
	@go mod download
	@echo "All package installed"

run:
	@go run ./main.go

# TODO currently not supported
watch:
	@air -c air.conf

build:
	swag init
	@go build -o ./build/main ./
