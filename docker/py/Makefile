ALL=py-debian \
	py-alpine \
	py-download

.PHONY: all
all: $(ALL)

.PHONY: run
run:
	env FLASK_APP=hello-world.py flask run -h 0.0.0.0 -p 80

.PHONY: $(ALL)
py-debian:
	time docker build -qt hello-world-$@ -f Dockerfile.$@ .
py-alpine:
	time docker build -qt hello-world-$@ -f Dockerfile.$@ .
py-download:
	time docker build -qt hello-world-$@ -f Dockerfile.$@ .
