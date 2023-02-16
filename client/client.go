package client

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks"
)

//client struct literal
// client := eks.Options{
// 	Region: region,
// 	// HTTPClient: ht,
// 	// Credentials: creds,
// 	// APIOptions: nil,
// 	// ClientLogMode: ,
// 	// DefaultsMode: ,
// 	// EndpointOptions: ,
// 	// HTTPSignerV4: nil,
// 	// EndpointOptions: ,
// 	// EndpointResolver: nil,
// 	// Logger: nil,
// 	// RetryMaxAttempts: ,
// 	// IdempotencyTokenProvider: nil,
// 	// Retryer: nil,
// 	// RetryMode: ,
// }

// better way to export Client
var Client *eks.Client

func CreateClient(cfg aws.Config) {

	Client = eks.NewFromConfig(cfg)
	//error handling
	//call create client?

	fmt.Print("client configured: ", Client)
}
