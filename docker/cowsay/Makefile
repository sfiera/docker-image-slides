ALL=bad \
	copy-first \
    install-first

.PHONY: all
all: $(ALL)

.PHONY: run
run:
	go run csaas.go

.PHONY: $(ALL)
bad:
	time docker build -qt csaas-$@ -f Dockerfile.$@ .
copy-first:
	time docker build -qt csaas-$@ -f Dockerfile.$@ .
install-first:
	time docker build -qt csaas-$@ -f Dockerfile.$@ .

rebuild-copy-first:
	docker image prune --all --force
	time docker build -qt csaas-copy-first -f Dockerfile.copy-first .
	sed -i '' s/cowsay/cowthink/ csaas.go
	time docker build -qt csaas-copy-first -f Dockerfile.copy-first .
	git checkout csaas.go

rebuild-install-first:
	docker image prune --all --force
	time docker build -qt csaas-install-first -f Dockerfile.install-first .
	sed -i '' s/cowsay/cowthink/ csaas.go
	time docker build -qt csaas-install-first -f Dockerfile.install-first .
	git checkout csaas.go
