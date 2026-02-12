package labels

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LabelsClient struct {
	Client *dataplane.Client
}

func NewLabelsClientUnconfigured() (*LabelsClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "labels", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LabelsClient: %+v", err)
	}

	return &LabelsClient{
		Client: client,
	}, nil
}

func (c *LabelsClient) LabelsClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewLabelsClientWithBaseURI(endpoint string) (*LabelsClient, error) {
	client, err := dataplane.NewClient(endpoint, "labels", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LabelsClient: %+v", err)
	}

	return &LabelsClient{
		Client: client,
	}, nil
}

func (c *LabelsClient) Clone(endpoint string) *LabelsClient {
	return &LabelsClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
