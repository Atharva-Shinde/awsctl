package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
)

type Handler struct {
	client       *eks.Client
	clusterInput *eks.CreateClusterInput
}

var cluster = Handler{
	clusterInput: &eks.CreateClusterInput{
		Name: &name,
		ResourcesVpcConfig: &types.VpcConfigRequest{
			EndpointPrivateAccess: nil,
			EndpointPublicAccess:  nil,
			PublicAccessCidrs:     nil,
			SecurityGroupIds:      nil,
			SubnetIds:             nil,
		},
		RoleArn:                 nil,
		ClientRequestToken:      nil,
		EncryptionConfig:        nil,
		KubernetesNetworkConfig: nil,
		Logging:                 nil,
		OutpostConfig:           nil,
		Tags:                    nil,
		Version:                 nil,
	},
}

func (h *Handler) Create() {
	// foo:= Handler{} //initialise empty struct?
	ctx := context.Background()
	result, err := h.client.CreateCluster(ctx, cluster.clusterInput)
	//TODO: error handling
	if err != nil {
		log.Fatal(err)
	}
	fmt.Sprintf("success:", result)

}
