DSN="root:@tcp(127.0.0.1:3306)/superindo?charset=utf8mb4&parseTime=True&loc=Local"
BINARY_NAME=superindo
ENV=development

## build: Build binary
build:
	@echo "Building back end..."
	go build -o ${BINARY_NAME} ./
	@echo "Binary built!"

## run: builds and runs the application
run: build
	@echo "Starting back end..."
	@env DSN=${DSN} ENV=${ENV} ./${BINARY_NAME} &
	@echo "Back end started!"

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
	@echo "Stopping back end..."
	@-pkill -SIGTERM -f "./${BINARY_NAME}"
	@echo "Stopped back end!"

## restart: stops and starts the running application
restart: stop start