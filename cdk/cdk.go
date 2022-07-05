package main

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type CdkStackProps struct {
	awscdk.StackProps
}

func NewCdkStack(scope constructs.Construct, id string, props *CdkStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here

	addLambda := awscdklambdagoalpha.NewGoFunction(stack, jsii.String("AddLambda"), &awscdklambdagoalpha.GoFunctionProps{
		Runtime:      awslambda.Runtime_GO_1_X(),
		Architecture: awslambda.Architecture_X86_64(),
		Entry:        jsii.String("../api/handlers/maths/add/post"),
		MemorySize:   jsii.Number(1024),
		Environment: &map[string]*string{
			"SENTENCE": jsii.String("HELLO WORLD"),
		},
		Bundling: &awscdklambdagoalpha.BundlingOptions{
			GoBuildFlags: &[]*string{jsii.String(fmt.Sprintf(`-ldflags "%s"`, "-s -w"))},
		},
	})

	notFound := awscdklambdagoalpha.NewGoFunction(stack, jsii.String("NotFoundLambda"), &awscdklambdagoalpha.GoFunctionProps{
		Runtime:      awslambda.Runtime_GO_1_X(),
		Architecture: awslambda.Architecture_X86_64(),
		Entry:        jsii.String("../api/handlers/maths/notfound"),
		MemorySize:   jsii.Number(1024),
		Bundling: &awscdklambdagoalpha.BundlingOptions{
			GoBuildFlags: &[]*string{jsii.String(fmt.Sprintf(`-ldflags "%s"`, "-s -w"))},
		},
	})

	api := awsapigateway.NewLambdaRestApi(stack, jsii.String("maths-api"), &awsapigateway.LambdaRestApiProps{
		Handler:        notFound,
		CloudWatchRole: jsii.Bool(false),
	})

	mathsRoot := api.Root()
	mathsAddEndpoint := mathsRoot.AddResource(jsii.String("add"), &awsapigateway.ResourceOptions{})
	mathsAddEndpoint.AddMethod(jsii.String("POST"), awsapigateway.NewLambdaIntegration(addLambda, &awsapigateway.LambdaIntegrationOptions{}), &awsapigateway.MethodOptions{
		AuthorizationType: awsapigateway.AuthorizationType_IAM,
	})

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewCdkStack(app, "maths-api-london", &CdkStackProps{
		awscdk.StackProps{
			Env: &awscdk.Environment{
				Region: jsii.String("eu-west-2"),
			},
		},
	})
	NewCdkStack(app, "maths-api-2", &CdkStackProps{
		awscdk.StackProps{
			Env: &awscdk.Environment{
				Region: jsii.String("eu-west-1"),
			},
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
