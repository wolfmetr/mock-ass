NAME = mock-ass
MOCK_ASS_DATA_DIR = $$(pwd)/data

run:
	MOCK_ASS_DATA_DIR=${MOCK_ASS_DATA_DIR} $(NAME)

install:
	go install ./cmd/mock-ass

test:
	go test -race $(shell go list ./...| grep -v vendor)

test-vendor:
	go test -race ./vendor/...
