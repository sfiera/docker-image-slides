ALL=go-debian \
	go-debian-2 \
	go-alpine \
	go-alpine-2 \
	go-minideb \
	go-download

.PHONY: all
all: $(ALL)

.PHONY: run
run:
	go run hello-world.go

.PHONY: $(ALL)
go-debian:
	time docker build -qt hello-world-$@ -f Dockerfile.$@ .
go-debian-2:
	time docker build -qt hello-world-$@ -f Dockerfile.$@ .
go-alpine:
	time docker build -qt hello-world-$@ -f Dockerfile.$@ .
go-alpine-2:
	time docker build -qt hello-world-$@ -f Dockerfile.$@ .
go-minideb:
	time docker build -qt hello-world-$@ -f Dockerfile.$@ .
go-download:
	time docker build -qt hello-world-$@ -f Dockerfile.$@ .

rebuild-alpine-2:
	docker image prune --all --force
	time docker build -qt hello-world-go-alpine-2 -f Dockerfile.go-alpine-2 .
	sed -i '' s/Hello/Goodbye/ hello-world.go
	time docker build -qt hello-world-go-alpine-2 -f Dockerfile.go-alpine-2 .
	git checkout hello-world.go

rebuild-download:
	docker image prune --all --force
	time docker build -qt hello-world-go-download -f Dockerfile.go-download .
	sed -i '' s/Hello/Goodbye/ hello-world.go
	time docker build -qt hello-world-go-download -f Dockerfile.go-download .
	git checkout hello-world.go
