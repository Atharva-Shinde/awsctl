package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createClusterCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c"},
	Run: func(cmd *cobra.Command, args []string) {
		err := worker()
		if err != nil {
			fmt.Print(err)
		}
	},
}

var (
	region                string
	name                  string
	clientRequestToken    string
	endpointPrivateAccess bool
	endpointPublicAccess  bool
	publicAccessCidrs     []string
	securityGroupIds      []string
	subnetIds             []string
)

func init() {
	createClusterCmd.Flags().StringVarP(&region, "region", "r", "", "Region for cluster")
	createClusterCmd.Flags().StringVarP(&name, "name", "n", "", "Name of cluster")
	createClusterCmd.Flags().StringVarP(&clientRequestToken, "clientRequestToken", "token", "", "ClientRequestToken")
	createClusterCmd.Flags().StringSliceVarP(&publicAccessCidrs, "publicAccessCidrs", "pubcidr", nil, "PublicAccessCidrs")
	createClusterCmd.Flags().StringSliceVarP(&securityGroupIds, "securityGroupIds", "secgrpid", nil, "SecurityGroupIds")
	createClusterCmd.Flags().StringSliceVarP(&subnetIds, "subnetIds", "subnetips", nil, "SubnetIds")
	createClusterCmd.Flags().BoolVarP(&endpointPrivateAccess, "endpointPrivateAccess", "prvtaccess", false, "EndpointPrivateAccess")
	createClusterCmd.Flags().BoolVarP(&endpointPublicAccess, "endpointPublicAccess", "pubaccess", false, "EndpointPublicAccess")
	createClusterCmd.MarkFlagsRequiredTogether("region", "name", "ClientRequestToken")
}
