.DEFAULT_GOAL := start

go-server-image-name := sample-ddd-app-server-image
mysql-container-name := sample-ddd-app-db
server-container-name := sample-ddd-app-server

.PHONY: start
start: build
	docker compose up -d

.PHONY: stop
stop:
	docker compose down

.PHONY: mysql
mysql:
	docker exec -it $(mysql-container-name) mysql -u root -ppassword sample

.PHONY: server
server:
	docker exec -it $(server-container-name) bash

.PHONY: test
test: build
	docker run --rm -v $(CURDIR):/go/src/work \
		$(go-server-image-name) \
		go test ./... -v

.PHONY: build
build:
	docker build -t $(go-server-image-name) .
