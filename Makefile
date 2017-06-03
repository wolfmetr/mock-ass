NAME = mock-ass
MOCK_ASS_DATA_DIR = $$(pwd)/data

TEST_ARGS :=

run:
	MOCK_ASS_DATA_DIR=${MOCK_ASS_DATA_DIR} $(NAME)

install:
	go install ./cmd/mock-ass

test:
	go test -race $(shell go list ./...| grep -v vendor) ${ARGS}

test-v:
	make test ARGS=-v

test-vendor:
	go test -race ./vendor/... ${ARGS}

test-vendor-v:
	make test-vendor ARGS=-v

itest:
	./itest.sh & make run

