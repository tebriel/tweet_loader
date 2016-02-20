PROJECT_NAME ?= tweet_loader
ES_URL ?= http://frodux.in:9200
FILENAME ?= tweets.csv

.PHONY: docker_build docker_run build run test

docker_build:
	docker build --no-cache -t cmoultrie/${PROJECT_NAME} .

docker_run:
	docker run --rm -it \
		-e CONSUMER_KEY \
		-e CONSUMER_SECRET \
		-e ACCESS_TOKEN \
		-e ACCESS_TOKEN_SECRET \
		-e ES_URL \
		-l SERVICE_NAME=${PROJECT_NAME} \
		-l SERVICE_TAGS=http \
		--dns 10.77.2.12 \
		-p 9090:8080 \
		cmoultrie/${PROJECT_NAME} \
		${PROJECT_NAME} -v=2 -logtostderr=true

build:
	go build

run: build
	tweet_loader -v=2 -logtostderr=true -f ${FILENAME} -es_url ${ES_URL}

test:
	go test -v ./...
