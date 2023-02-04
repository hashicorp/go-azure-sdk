package v2021_04_01_preview

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/aadb2c/2021-04-01-preview/tenants"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Tenants *tenants.TenantsClient
}

func NewClientWithBaseURI(api environments.Api, configureAuthFunc func(c *resourcemanager.Client)) (*Client, error) {
	tenantsClient, err := tenants.NewTenantsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for Tenants: %+v", err)
	}
	configureAuthFunc(tenantsClient.Client)

	return &Client{
		Tenants: tenantsClient,
	}, nil
}
