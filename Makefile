TOKEN ?=

.PHONY: build run clean stop

NAME=golack
VERSION=1

build:
	- $(MAKE) clean
	docker build -t ${NAME}:${VERSION} .

run:
	docker run -d --env TOKEN=$(TOKEN) --name ${NAME} ${NAME}:${VERSION}

contener=`docker ps -a -q`
image=`docker images | awk '/^<none>/ { print $$3 }'`

clean: stop
	@if [ "$(image)" != "" ] ; then \
		docker rmi -f $(image); \
		fi
	@if [ "$(contener)" != "" ] ; then \
		docker rm -f $(contener); \
		fi

stop:
	docker rm -f $(NAME)
