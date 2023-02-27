// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2021_07_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/purview/2021-07-01/account"
	"github.com/hashicorp/go-azure-sdk/resource-manager/purview/2021-07-01/defaultaccount"
	"github.com/hashicorp/go-azure-sdk/resource-manager/purview/2021-07-01/privateendpointconnection"
	"github.com/hashicorp/go-azure-sdk/resource-manager/purview/2021-07-01/privatelinkresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/purview/2021-07-01/provider"
)

type Client struct {
	Account                   *account.AccountClient
	DefaultAccount            *defaultaccount.DefaultAccountClient
	PrivateEndpointConnection *privateendpointconnection.PrivateEndpointConnectionClient
	PrivateLinkResource       *privatelinkresource.PrivateLinkResourceClient
	Provider                  *provider.ProviderClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	accountClient := account.NewAccountClientWithBaseURI(endpoint)
	configureAuthFunc(&accountClient.Client)

	defaultAccountClient := defaultaccount.NewDefaultAccountClientWithBaseURI(endpoint)
	configureAuthFunc(&defaultAccountClient.Client)

	privateEndpointConnectionClient := privateendpointconnection.NewPrivateEndpointConnectionClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionClient.Client)

	privateLinkResourceClient := privatelinkresource.NewPrivateLinkResourceClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkResourceClient.Client)

	providerClient := provider.NewProviderClientWithBaseURI(endpoint)
	configureAuthFunc(&providerClient.Client)

	return Client{
		Account:                   &accountClient,
		DefaultAccount:            &defaultAccountClient,
		PrivateEndpointConnection: &privateEndpointConnectionClient,
		PrivateLinkResource:       &privateLinkResourceClient,
		Provider:                  &providerClient,
	}
}
