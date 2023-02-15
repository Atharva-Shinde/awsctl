package authenticator

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type Credentials struct {
	creds aws.Credentials
}

type Config struct {
	cfg aws.Config
}

func GetCredentials() error {
	scanner := bufio.NewReader(os.Stdin)
	fmt.Print("Enter credentials\n accesskeyid: ")
	acceskeyid, err := scanner.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("secretaccesskey: ")
	secretaccesskey, err := scanner.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	sessiontoken, err := scanner.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	setCredentials(acceskeyid, secretaccesskey, sessiontoken)

	return err
}

func setCredentials(acceskeyid, secretaccesskey, sessiontoken string) {
	var cr = Credentials{
		creds: aws.Credentials{
			AccessKeyID:     acceskeyid,
			SecretAccessKey: secretaccesskey,
			SessionToken:    sessiontoken,
			// CanExpire:,
			// Source: ,
			// Expires: ,
		},
	}
	cr.CredValidate()
}

func (c *Credentials) CredValidate() {

	switch {
	case c.creds.Expired():
		log.Fatal("creds expired")
	case c.creds.HasKeys():
		log.Fatal("need: accessid, secretid")
	case c.creds.Expired() && c.creds.HasKeys():
		// call createclient()  ?

	}

}

// func getConfig() {
// 	getCredentials()
// }

// BearerAuthToken struct literal
// var token = bearer.Token{
// 	Value:     bearerAuthToken,
// 	CanExpire: false,
// 	Expires: ,
// }
// func (b *Config) TokenValidate() {
// 	switch {
// 	case token.Expired(expires):

// 	}
// }
