.PHONY: slides.pdf
slides.pdf:
	docker run --rm --net=host -v `pwd`:/slides astefanutti/decktape -s 1920x1080 http://host.docker.internal:8000 slides.pdf
