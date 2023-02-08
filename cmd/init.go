package cmd

import (
	"time"

	"github.com/spf13/cobra"
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

type Creds struct {
	Accessid        string    //creds
	SecretAccessKey string    //creds
	SessionToken    string    //creds
	Canexpire       bool      //creds
	Source          string    //creds
	Expires         time.Time //creds
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

var (
	region          string  //config, eks.options
	bearerAuthToken string  //config
	ClusterInfo     Cluster //cluster
	Credentials     Creds   //credentials
)

var createClusterCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c"},
	Run: func(cmd *cobra.Command, args []string) {

		//-client
		// client := &Client{
		// 	opt: eks.Options{
		// 		Region: region,
		// 		HTTPClient: ht,
		// 		Credentials: creds,
		// 		APIOptions: nil,
		// 		ClientLogMode: ,
		// 		DefaultsMode: ,
		// 		EndpointOptions: ,
		// 		HTTPSignerV4: nil,
		// 		EndpointOptions: ,
		// 		EndpointResolver: nil,
		// 		Logger: nil,
		// 		RetryMaxAttempts: ,
		// 		IdempotencyTokenProvider: nil,
		// 		Retryer: nil,
		// 		RetryMode: ,
		// 	},
		//  }
		// client.CreateClient()
	},
}

func init() {
	createClusterCmd.Flags().StringVarP(&region, "region", "r", "", "Region for cluster")                                                                           //region
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
