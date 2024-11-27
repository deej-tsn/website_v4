REGISTRY_HOST := deejtsn/private
IMAGE := website_v4

run:
	templ generate && go build -o ./tmp/main cmd/website_v4/main.go

build-arm:
	docker build -t ${REGISTRY_HOST}:${IMAGE}_arm --platform=linux/arm64 .

build-amd64:
	docker build -t ${REGISTRY_HOST}:${IMAGE}_amd64 --platform=linux/amd64 .

build-push-amd64:
	docker build -t ${REGISTRY_HOST}:${IMAGE}_amd64 --platform=linux/amd64 . && docker push ${REGISTRY_HOST}:amd64