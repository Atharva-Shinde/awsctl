package cmd

import (
	"github.com/spf13/cobra"
	"main.go/authenticator"
)

var (
	Region          string  //config, eks.options
	bearerAuthToken string  //config
	ClusterInfo     Cluster //cluster
)

type Cluster struct {
	Name                    string
	ResourcesVpcConfig      Rvc
	RoleArn                 string
	ClientRequestToken      string
	KubernetesNetworkConfig Knc
	Logging                 Lg
	Tags                    map[string]string
	Version                 string
}

type Rvc struct {
	EndpointPrivateAccess bool
	EndpointPublicAccess  bool
	PublicAccessCidrs     []string
	SecurityGroupIds      []string
	SubnetIds             []string
}

type Knc struct {
	IpFamily        string
	ServiceIpv4Cidr string
}

type Lg struct {
	Enable bool
	// TODO: LogType manage
	// LogTypeApi               LogType = "api"
	// LogTypeAudit             LogType = "audit"
	// LogTypeAuthenticator     LogType = "authenticator"
	// LogTypeControllerManager LogType = "controllerManager"
	// LogTypeScheduler         LogType = "scheduler"
}

var createClusterCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c"},
	Run: func(cmd *cobra.Command, args []string) {
		authenticator.GetCredentials()
	},
}

func init() {
	createClusterCmd.Flags().StringVarP(&Region, "region", "r", "", "Region for cluster")                                                                           //region
	createClusterCmd.Flags().StringVarP(&ClusterInfo.Name, "name", "n", "", "Name of cluster")                                                                      //name
	createClusterCmd.Flags().StringVarP(&ClusterInfo.ClientRequestToken, "clientRequestToken", "token", "", "ClientRequestToken")                                   //clientRequestToken
	createClusterCmd.Flags().StringSliceVarP(&ClusterInfo.ResourcesVpcConfig.PublicAccessCidrs, "publicAccessCidrs", "pubcidr", nil, "PublicAccessCidrs")           //publicAccessCidrs
	createClusterCmd.Flags().StringSliceVarP(&ClusterInfo.ResourcesVpcConfig.SecurityGroupIds, "securityGroupIds", "secgrpid", nil, "SecurityGroupIds")             //securityGroupIds
	createClusterCmd.Flags().StringSliceVarP(&ClusterInfo.ResourcesVpcConfig.SubnetIds, "subnetIds", "subnetids", nil, "SubnetIds")                                 //subnetids
	createClusterCmd.Flags().BoolVarP(&ClusterInfo.ResourcesVpcConfig.EndpointPrivateAccess, "endpointPrivateAccess", "prvtaccess", false, "EndpointPrivateAccess") //prvtaccess
	createClusterCmd.Flags().BoolVarP(&ClusterInfo.ResourcesVpcConfig.EndpointPublicAccess, "endpointPublicAccess", "pubaccess", false, "EndpointPublicAccess")     //pubaccess
	createClusterCmd.Flags().StringVarP(&ClusterInfo.RoleArn, "rolearn", "role", "", "RoleArn")                                                                     //rolearn
	createClusterCmd.Flags().StringVarP(&ClusterInfo.Version, "version", "v", "", "Version")                                                                        //version
	createClusterCmd.Flags().BoolVarP(&ClusterInfo.Logging.Enable, "logging", "log", false, "Logging")                                                              //logging
	createClusterCmd.Flags().StringVarP(&ClusterInfo.KubernetesNetworkConfig.IpFamily, "ip", "ip", "", "KuberneteNetworkIp")                                        //ipfamily
	createClusterCmd.Flags().StringVarP(&ClusterInfo.KubernetesNetworkConfig.ServiceIpv4Cidr, "srvcidr", "cidr", "", "ServiceIpv4Cidr")                             //serviceIpv4Cidr
	//TODO: tag
	createClusterCmd.MarkFlagsRequiredTogether("region", "name", "ClientRequestToken")
}
