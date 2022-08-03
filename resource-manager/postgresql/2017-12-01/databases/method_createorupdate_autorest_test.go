package databases_test

import (
	"context"
	"os"
	"testing"

	"github.com/hashicorp/go-azure-sdk/auth"
	"github.com/hashicorp/go-azure-sdk/environments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2017-12-01/databases"
)

func TestCreate(t *testing.T) {
	client := databases.NewDatabasesClientWithBaseURI(string(environments.Global.ResourceManager.Endpoint))
	client.Client2.Client.Authorizer = (&auth.ClientCredentialsConfig{
		Environment:  environments.Global,
		TenantID:     os.Getenv("ARM_TENANT_ID"),
		ClientID:     os.Getenv("ARM_CLIENT_ID"),
		ClientSecret: os.Getenv("ARM_CLIENT_SECRET"),
		Scopes:       []string{environments.Global.ResourceManager.DefaultScope()},
		TokenVersion: auth.TokenVersion2,
	}).TokenSource(context.TODO(), auth.ClientCredentialsSecretType)

	id := databases.NewDatabaseID(os.Getenv("ARM_SUBSCRIPTION_ID"), "aaacctest-tbamford", "aabtest-tbamford", "aaa1")
	charset := "UTF8"
	collation := "English_United States.1252"
	input := databases.Database{
		Properties: &databases.DatabaseProperties{
			Charset:   &charset,
			Collation: &collation,
		},
	}

	err := client.CreateOrUpdateThenPoll(context.TODO(), id, input)
	if err != nil {
		t.Fatal(err)
	}
}
