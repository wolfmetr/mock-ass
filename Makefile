NAME = mock-ass
MOCK_ASS_DATA_DIR = $$(pwd)/data

run:
	MOCK_ASS_DATA_DIR=${MOCK_ASS_DATA_DIR} $(NAME) -color

install:
	go install ./cmd/mock-ass
