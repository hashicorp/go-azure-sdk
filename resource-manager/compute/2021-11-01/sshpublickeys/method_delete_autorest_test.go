package sshpublickeys_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/go-azure-sdk/auth"
	"github.com/hashicorp/go-azure-sdk/environments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/sshpublickeys"
)

func TestDelete(t *testing.T) {
	client := sshpublickeys.NewSshPublicKeysClientWithBaseURI(string(environments.Global.ResourceManager.Endpoint))
	client.Client2.Client.Authorizer = (&auth.ClientCredentialsConfig{
		Environment:  environments.Global,
		TenantID:     os.Getenv("ARM_TENANT_ID"),
		ClientID:     os.Getenv("ARM_CLIENT_ID"),
		ClientSecret: os.Getenv("ARM_CLIENT_SECRET"),
		Scopes:       []string{environments.Global.ResourceManager.DefaultScope()},
		TokenVersion: auth.TokenVersion2,
	}).TokenSource(context.TODO(), auth.ClientCredentialsSecretType)

	id := sshpublickeys.NewSshPublicKeyID(os.Getenv("ARM_SUBSCRIPTION_ID"), "aaacctest-tbamford", "aaa3")

	res, err := client.Delete(context.TODO(), id)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v\n", res)
}
