project_name ?= tweet_loader

.PHONY: build docker_run run test

build:
	docker build --no-cache -t cmoultrie/${project_name} .

docker_run:
	docker run --rm -it \
		-e CONSUMER_KEY \
		-e CONSUMER_SECRET \
		-e ACCESS_TOKEN \
		-e ACCESS_TOKEN_SECRET \
		-e ES_URL \
		-l SERVICE_NAME=${project_name} \
		-l SERVICE_TAGS=http \
		--dns 10.77.2.12 \
		-p 9090:8080 \
		cmoultrie/${project_name} \
		${project_name} -v=2 -logtostderr=true

run:
	go run *.go -v=2 -logtostderr=true

test:
	go test ./...
