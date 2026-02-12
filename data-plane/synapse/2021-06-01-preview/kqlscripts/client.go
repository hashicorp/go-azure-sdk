package kqlscripts

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KqlScriptsClient struct {
	Client *dataplane.Client
}

func NewKqlScriptsClientUnconfigured() (*KqlScriptsClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "kqlscripts", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating KqlScriptsClient: %+v", err)
	}

	return &KqlScriptsClient{
		Client: client,
	}, nil
}

func (c *KqlScriptsClient) KqlScriptsClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewKqlScriptsClientWithBaseURI(endpoint string) (*KqlScriptsClient, error) {
	client, err := dataplane.NewClient(endpoint, "kqlscripts", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating KqlScriptsClient: %+v", err)
	}

	return &KqlScriptsClient{
		Client: client,
	}, nil
}

func (c *KqlScriptsClient) Clone(endpoint string) *KqlScriptsClient {
	return &KqlScriptsClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
