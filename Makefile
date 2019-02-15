TOKEN ?=

.PHONY: build run clean stop

NAME=golack
VERSION=1

build:
	- $(MAKE) stop
	docker build -t ${NAME}:${VERSION} .

run:
	docker run -d --env TOKEN=$(TOKEN) --name ${NAME} ${NAME}:${VERSION}

stop:
	docker rm -f $(NAME)
