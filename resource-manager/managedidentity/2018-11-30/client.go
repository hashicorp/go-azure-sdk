package v2018_11_30

import (
	"github.com/hashicorp/go-azure-sdk/resource-manager/managedidentity/2018-11-30/managedidentity"
)

type Client struct {
	ManagedIdentity *managedidentity.ManagedIdentityClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	managedIdentityClient := managedidentity.NewManagedIdentityClientWithBaseURI(endpoint)
	configureAuthFunc(&managedIdentityClient.Client)

	return Client{
		ManagedIdentity: &managedIdentityClient,
	}
}
