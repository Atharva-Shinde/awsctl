package cmd

import (
	// "github.com/aws/aws-sdk-go-v2/aws"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/eks"
)

type Configuration struct {
	config      config.Config
	credentials aws.Credentials
}

func (config *Configuration) worker() {

	// c := Configuration{
	// 	config: aws.Config{
	// 		Region: "india",
	// 	},
	// 	credentials: aws.Credentials{
	// 		AccessKeyID:     "kjsjlksjl",
	// 		SecretAccessKey: "lsjjlsd",
	// 	},
	// }

	x := eks.New(eks.Options{Credentials: aws.NewConfig().Credentials})
	ctx := context.Background()
	params := &eks.CreateClusterInput{
		Name: aws.String("mycluster"),
	}
	output, err := x.CreateCluster(ctx, params)
	if err != nil {
		fmt.Print("error")
	}
	fmt.Print(output)

}
