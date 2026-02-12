package policy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyClient struct {
	Client *dataplane.Client
}

func NewPolicyClientUnconfigured() (*PolicyClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "policy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyClient: %+v", err)
	}

	return &PolicyClient{
		Client: client,
	}, nil
}

func (c *PolicyClient) PolicyClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewPolicyClientWithBaseURI(endpoint string) (*PolicyClient, error) {
	client, err := dataplane.NewClient(endpoint, "policy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyClient: %+v", err)
	}

	return &PolicyClient{
		Client: client,
	}, nil
}

func (c *PolicyClient) Clone(endpoint string) *PolicyClient {
	return &PolicyClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
