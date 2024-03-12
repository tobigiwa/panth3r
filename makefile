BINARY_NAME=Panth3r

run:
	@echo "Running server"
	@./bin/${BINARY_NAME}

Productionbuild:
	@echo "Building binary for production"
	@go build -mod vendor -tags netgo -ldflags '-s -w' -o bin/${BINARY_NAME} cmd/Panth3r/main.go

Devbuild:
	@echo "Building binary for development"
	@go build -mod vendor -o bin/${BINARY_NAME} cmd/Panth3r/main.go


swagger:
	@echo "Generating swagger"
	@swag init -d cmd/panth3r/,http/
	@swag fmt --exclude internal/,templ/,env/

dev: swagger Devbuild run

prod: swagger Productionbuild run

