package openidmetadatadiscovery

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenIDMetadataDiscoveryClient struct {
	Client *dataplane.Client
}

func NewOpenIDMetadataDiscoveryClientUnconfigured() (*OpenIDMetadataDiscoveryClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "openidmetadatadiscovery", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OpenIDMetadataDiscoveryClient: %+v", err)
	}

	return &OpenIDMetadataDiscoveryClient{
		Client: client,
	}, nil
}

func (c *OpenIDMetadataDiscoveryClient) OpenIDMetadataDiscoveryClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewOpenIDMetadataDiscoveryClientWithBaseURI(endpoint string) (*OpenIDMetadataDiscoveryClient, error) {
	client, err := dataplane.NewClient(endpoint, "openidmetadatadiscovery", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OpenIDMetadataDiscoveryClient: %+v", err)
	}

	return &OpenIDMetadataDiscoveryClient{
		Client: client,
	}, nil
}

func (c *OpenIDMetadataDiscoveryClient) Clone(endpoint string) *OpenIDMetadataDiscoveryClient {
	return &OpenIDMetadataDiscoveryClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
