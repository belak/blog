build: build-latest build-draft

build-latest:
	docker build -t belak/blog:latest .

build-draft:
	docker build -f Dockerfile.beta -t belak/blog:draft .

push: push-latest push-draft

push-latest: build-latest
	docker push belak/blog:latest

push-draft: build-draft
	docker push belak/blog:draft
