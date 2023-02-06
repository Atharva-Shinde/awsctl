package cmd

import (
	"github.com/spf13/cobra"
)

var (
	region                string   //cluster
	name                  string   //cluster
	clientRequestToken    string   //cluster
	endpointPrivateAccess bool     //cluster
	endpointPublicAccess  bool     //cluster
	publicAccessCidrs     []string //cluster
	securityGroupIds      []string //cluster
	subnetIds             []string //cluster
	// accessid              string    //creds
	// secretAccessKey       string    //creds
	// sessionToken          string    //creds
	// canexpire             bool      //creds
	// source                string    //creds
	// expires               time.Time //creds
	// BearerAuthToken       string    //config
)

var createClusterCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c"},
	Run: func(cmd *cobra.Command, args []string) {

		//-credentials TODO: way to provide credentials
		// creds := credentials{
		// 	cr: aws.Credentials{
		// 		AccessKeyID:     accessid,
		// 		SecretAccessKey: secretAccessKey,
		// 		SessionToken:    sessionToken,
		// 		CanExpire:       canexpire,
		// 		// Source: ,
		// 		// Expires: ,
		// 	},
		// }
		// creds.CredValidate()

		//-client
		// client := &Client{}
		// client.CreateClient()
	},
}

func init() {
	createClusterCmd.Flags().StringVarP(&region, "region", "r", "", "Region for cluster")                                            //region
	createClusterCmd.Flags().StringVarP(&name, "name", "n", "", "Name of cluster")                                                   //name
	createClusterCmd.Flags().StringVarP(&clientRequestToken, "clientRequestToken", "token", "", "ClientRequestToken")                //clientRequestToken
	createClusterCmd.Flags().StringSliceVarP(&publicAccessCidrs, "publicAccessCidrs", "pubcidr", nil, "PublicAccessCidrs")           //publicAccessCidrs
	createClusterCmd.Flags().StringSliceVarP(&securityGroupIds, "securityGroupIds", "secgrpid", nil, "SecurityGroupIds")             //securityGroupIds
	createClusterCmd.Flags().StringSliceVarP(&subnetIds, "subnetIds", "subnetids", nil, "SubnetIds")                                 //subnetids
	createClusterCmd.Flags().BoolVarP(&endpointPrivateAccess, "endpointPrivateAccess", "prvtaccess", false, "EndpointPrivateAccess") //prvtaccess
	createClusterCmd.Flags().BoolVarP(&endpointPublicAccess, "endpointPublicAccess", "pubaccess", false, "EndpointPublicAccess")     //pubaccess
	createClusterCmd.MarkFlagsRequiredTogether("region", "name", "ClientRequestToken")
}
