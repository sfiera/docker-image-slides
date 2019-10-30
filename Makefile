.PHONY: help
help:
	@ echo "usage: make serve -- serve and reload on localhost:8000"
	@ echo "       make docs  -- export slides to docs/"
	@ echo "       make pdf   -- export slides to slides.pdf"

.PHONY: serve
serve:
	hovercraft slides.rst

.PHONY: docs
docs:
	hovercraft slides.rst docs/

.PHONY: pdf
pdf:
	docker run --rm -v `pwd`:/slides \
		astefanutti/decktape -s 1920x1080 \
		http://host.docker.internal:8000 \
		slides.pdf
