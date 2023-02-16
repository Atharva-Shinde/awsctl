package config

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"main.go/client"
	"main.go/cmd"
)

// type Config struct {
// 	Region string
// 	Credentials
// 	BearerAuthToken             Bat
// 	HTTPClient                  Httpc
// 	EndpointResolver            Er
// 	EndpointResolverWithOptions Eropt
// 	RetryMaxAttempts            Rma
// 	RetryMode                   Rm
// }

func SetConfigHandler(cr aws.Credentials) {
	creds := []string{cr.AccessKeyID, cr.SecretAccessKey, cr.SessionToken}
	setConfig(creds, cmd.Region)
}

func setConfig(creds []string, region string) {
	//TODO: with no region specified
	//config.LoadDefaultConfig(context.TODO(),config.WithDefaultRegion)

	// TODO: ConfigWRegion standalone variable in the package
	ConfigWRegion, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region), config.WithSharedCredentialsFiles(creds))
	if err != nil {
		//improve error handling
		log.Fatal(err)
	}
	fmt.Print("config set:", ConfigWRegion)
	client.CreateClient(ConfigWRegion)
}

//TODO getConfig
