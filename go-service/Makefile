init:
	echo "Init"

build-docker-image:
	docker build -t go-service -f Dockerfile .

push:
	docker tag go-service gcr.io/golden-tenure-231009/go-service:2.0
	docker push gcr.io/golden-tenure-231009/go-service:2.0

all: init  build-docker-image push
