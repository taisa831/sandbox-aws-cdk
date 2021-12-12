# Welcome to your CDK Go project!

This is a blank project for Go development with CDK.

**NOTICE**: Go support is still in Developer Preview. This implies that APIs may
change while we address early feedback from the community. We would love to hear
about your experience through GitHub issues.

## Useful commands

 * `cdk deploy`      deploy this stack to your default AWS account/region
 * `cdk diff`        compare deployed stack with current state
 * `cdk synth`       emits the synthesized CloudFormation template
 * `go test`         run unit tests


## Build

```
$ GOARCH=amd64 GOOS=linux go build -o bin/main lambda/main.go
```

## Deploy

```
$ cdk deploy --profile p
```

## Run

```
curl -X POST -H "Content-Type: application/json" -d '{"name":"world!"}' https://hd4t4cahyl.execute-api.ap-northeast-1.amazonaws.com/prod/hello
```