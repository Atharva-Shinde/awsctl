package createcluster

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"main.go/client"
	"main.go/cmd"
)

type Handler struct {
	clientAgent  *eks.Client
	clusterInput *eks.CreateClusterInput
}

var cluster = Handler{
	clientAgent: client.Client,
	clusterInput: &eks.CreateClusterInput{
		Name: &cmd.ClusterInfo.Name,
		ResourcesVpcConfig: &types.VpcConfigRequest{
			EndpointPrivateAccess: &cmd.ClusterInfo.ResourcesVpcConfig.EndpointPrivateAccess,
			EndpointPublicAccess:  &cmd.ClusterInfo.ResourcesVpcConfig.EndpointPublicAccess,
			PublicAccessCidrs:     cmd.ClusterInfo.ResourcesVpcConfig.PublicAccessCidrs,
			SecurityGroupIds:      cmd.ClusterInfo.ResourcesVpcConfig.SecurityGroupIds,
			SubnetIds:             cmd.ClusterInfo.ResourcesVpcConfig.SubnetIds,
		},
		RoleArn:            &cmd.ClusterInfo.RoleArn,
		ClientRequestToken: &cmd.ClusterInfo.ClientRequestToken,
		EncryptionConfig:   nil,
		KubernetesNetworkConfig: &types.KubernetesNetworkConfigRequest{
			IpFamily:        types.IpFamily(cmd.ClusterInfo.KubernetesNetworkConfig.IpFamily),
			ServiceIpv4Cidr: &cmd.ClusterInfo.KubernetesNetworkConfig.ServiceIpv4Cidr,
		},
		Logging: &types.Logging{
			ClusterLogging: []types.LogSetup{
				{
					Enabled: &cmd.ClusterInfo.Logging.Enable,
				},
			},
		},
		OutpostConfig: nil,
		Tags:          cmd.ClusterInfo.Tags,
		Version:       &cmd.ClusterInfo.Version,
	},
}

func (h *Handler) Create() {
	// foo:= Handler{} //initialise empty struct?
	ctx := context.Background()
	if cluster.clusterInput.ClientRequestToken == nil {
		fmt.Print("no token provided, generating idempotent token")
	}
	result, err := cluster.clientAgent.CreateCluster(ctx, cluster.clusterInput)
	if err != nil {
		HandleError(result)
	}
	fmt.Println("success: \ncluster info:", result.Cluster)

}

// TODO: error handling
func HandleError(result *eks.CreateClusterOutput) error {
	return nil
}
