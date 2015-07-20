NAME = mock-ass

all: build

build:
	go build -o $(NAME)

build-restore: restore build

run: all
	./$(NAME)

run-dev:
	go run main.go

restore:
	godep restore


