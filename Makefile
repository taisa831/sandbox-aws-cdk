build:
	GOARCH=amd64 GOOS=linux go build -o handler/main handler/main.go

deploy:
	cdk deploy --profile g1

rebuild:build deploy
