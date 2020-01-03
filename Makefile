
GO = docker run --rm -e GOOS=linux -e GOARCH=amd64 -v $(PWD):/usr/src/app -w /usr/src/app golang:1.13 go
R = docker run --rm --name tbot --network dev -e APP_ENV=dev -e GOOS=linux -e GOARCH=amd64 -v $(PWD):/usr/src/app -w /usr/src/app golang:1.13 go
LAMBDA_VERSION = 7

.PHONY: clean test build lambda_build lambda_artifact lambda_deploy deploy

build: clean test lambda_build lambda_artifact

run:
	$(R) run ./cmd/http.go

test:
	@echo "Testing..."
	$(GO) test cmd/lambda.go

lambda_build:
	@echo "Building..."
	$(GO) build -o build/out/app cmd/lambda.go

lambda_artifact:
	zip -j build/out/main.zip build/out/app

deploy:
	@echo "Deploying..."
	aws lambda update-function-code --function-name TBotLearn \
    --zip-file "fileb://build/out/main.zip"

clean:
	rm build/out/app || true
	rm build/out/main.zip || true
