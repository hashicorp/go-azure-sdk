package v2020_01_01

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2020-01-01/serverkeys"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ServerKeys *serverkeys.ServerKeysClient
}

func NewClientWithBaseURI(api environments.Api, configureAuthFunc func(c *resourcemanager.Client)) (*Client, error) {
	serverKeysClient, err := serverkeys.NewServerKeysClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for ServerKeys: %+v", err)
	}
	configureAuthFunc(serverKeysClient.Client)

	return &Client{
		ServerKeys: serverKeysClient,
	}, nil
}
