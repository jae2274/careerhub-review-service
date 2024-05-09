BINARY_NAME=myapp
CODE_DIR=./careerhub/review_service
CONTAINER_IMAGE_NAME=careerhub-review-service

# include test.env

## build: Build binary
build:
	@echo "Building..."
	@go build -ldflags="-s -w" -o ${BINARY_NAME} ${CODE_DIR}
	@echo "Built!"

image_build:
	@echo "Building..."
	@docker build -t ${CONTAINER_IMAGE_NAME}:latest .
	@echo "Built!"

## run: builds and runs the application
run: build
	@echo "Starting..."
	@env ./${BINARY_NAME} 
	@echo "Started!"

## clean: runs go clean and deletes binaries
clean:
	@echo "Cleaning..."
	@go clean
	@rm ${BINARY_NAME}
	@echo "Cleaned!"

## start: an alias to run
start: run

## stop: stops the running application
stop:
	@echo "Stopping..."
	@-pkill -SIGTERM -f "./${BINARY_NAME}"
	@echo "Stopped!"

## restart: stops and starts the application
restart: stop start

proto:
	@protoc careerhub/review_service/crawler/crawler_grpc/*.proto  --go_out=. --go-grpc_out=. --go-grpc_opt=paths=source_relative  --go_opt=paths=source_relative  --proto_path=.
	@protoc careerhub/review_service/provider/provider_grpc/*.proto  --go_out=. --go-grpc_out=. --go-grpc_opt=paths=source_relative  --go_opt=paths=source_relative  --proto_path=.

## test: runs all tests
test:	
	@echo "Testing..."
	@env go test -p 1 -timeout 60s ./test/...
	

