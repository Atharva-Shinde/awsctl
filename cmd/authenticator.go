package cmd

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type credentials struct {
	cr aws.Credentials
}

type config struct {
	cfg aws.Config
}

// Credentials struct literal
var creds = aws.Credentials{
	AccessKeyID:     accessid,
	SecretAccessKey: secretAccessKey,
	SessionToken:    sessionToken,
	CanExpire:       canexpire,
}

func (c *credentials) CredValidate() {
	switch {
	case creds.Expired():
		log.Fatal("creds expired")
	case !creds.HasKeys():
		log.Fatal("need: accessid, secretid")
	case !creds.Expired() && creds.HasKeys():
		// TODO: createcluster()
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
