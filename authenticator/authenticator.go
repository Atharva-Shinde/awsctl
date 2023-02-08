package authenticator

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"main.go/cmd"
)

type Credentials struct {
	creds aws.Credentials
}

type config struct {
	cfg aws.Config
}

// -credentials TODO: way to provide credentials
var cr = Credentials{
	creds: aws.Credentials{
		AccessKeyID:     cmd.Accessid,
		SecretAccessKey: cmd.SecretAccessKey,
		SessionToken:    cmd.SessionToken,
		CanExpire:       cmd.Canexpire,
		// Source: ,
		// Expires: ,
	},
}

func (c *Credentials) CredValidate() {
	switch {
	case cr.creds.Expired():
		log.Fatal("creds expired")
	case !cr.creds.HasKeys():
		log.Fatal("need: accessid, secretid")
	case !cr.creds.Expired() && cr.creds.HasKeys():
		// call createcluster()  ?

	}

}

// BearerAuthToken struct literal
// var token = bearer.Token{
// 	Value:     bearerAuthToken,
// 	CanExpire: false,
// 	Expires: ,
// }
// func (b *config) TokenValidate() {
// 	switch {
// 	case token.Expired(expires):

// 	}
// }
