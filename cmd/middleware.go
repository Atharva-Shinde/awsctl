package cmd

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
)

type Configuration struct {
	cfg aws.Config
}

func worker() {
	cf := Configuration{
		cfg: aws.Config{
			Region:      "india",
			Credentials: aws.NewConfig().Credentials,
		},
	}

	x := eks.NewFromConfig(cf.cfg)
	clusterinfo := eks.CreateClusterInput{
		Name: aws.String("dlkjdl"),
		ResourcesVpcConfig: &types.VpcConfigRequest{
			EndpointPrivateAccess: aws.Bool(false),
			EndpointPublicAccess:  aws.Bool(false),
			PublicAccessCidrs:     aws.ToStringSlice([]*string{aws.String("dfjk")}),
			SecurityGroupIds:      aws.ToStringSlice([]*string{aws.String("oirio")}),
			SubnetIds:             aws.ToStringSlice([]*string{aws.String("aklkld")}),
		},
		RoleArn:                 aws.String("jldkd"),
		ClientRequestToken:      aws.String("lkjddl"),
		EncryptionConfig:        []types.EncryptionConfig{},
		KubernetesNetworkConfig: &types.KubernetesNetworkConfigRequest{IpFamily: "lfdjk", ServiceIpv4Cidr: aws.String("jlkd")},
		Logging:                 &types.Logging{ClusterLogging: []types.LogSetup{}},
		Version:                 aws.String("132"),
	}

	ctx := context.Background()
	b, err := x.CreateCluster(ctx, &clusterinfo)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(b)

}
