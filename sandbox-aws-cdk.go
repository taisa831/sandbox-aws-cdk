package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type SandboxAwsCdkStackProps struct {
	awscdk.StackProps
}

func NewSandboxAwsCdkStack(scope constructs.Construct, id string, props *SandboxAwsCdkStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	lambdaFn := awslambda.NewFunction(stack, jsii.String("my-lambda-aws-cdk"), &awslambda.FunctionProps{
		FunctionName: jsii.String("my-lambda-aws-cdk-func"),
		Runtime:      awslambda.Runtime_GO_1_X(),
		Code:         awslambda.AssetCode_FromAsset(jsii.String("handler"), &awss3assets.AssetOptions{}),
		Handler:      jsii.String("main"),
	})

	apiGW := awsapigateway.NewRestApi(stack, jsii.String("my-api-gw-aws-cdk"), nil)
	apiGW.Root().
		AddResource(jsii.String("hello"), nil).
		AddMethod(jsii.String("POST"), awsapigateway.NewLambdaIntegration(lambdaFn, nil), nil)

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewSandboxAwsCdkStack(app, "SandboxAwsCdkStack", &SandboxAwsCdkStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
