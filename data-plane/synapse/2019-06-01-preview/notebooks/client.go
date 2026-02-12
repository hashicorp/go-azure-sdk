package notebooks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotebooksClient struct {
	Client *dataplane.Client
}

func NewNotebooksClientUnconfigured() (*NotebooksClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "notebooks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating NotebooksClient: %+v", err)
	}

	return &NotebooksClient{
		Client: client,
	}, nil
}

func (c *NotebooksClient) NotebooksClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewNotebooksClientWithBaseURI(endpoint string) (*NotebooksClient, error) {
	client, err := dataplane.NewClient(endpoint, "notebooks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating NotebooksClient: %+v", err)
	}

	return &NotebooksClient{
		Client: client,
	}, nil
}

func (c *NotebooksClient) Clone(endpoint string) *NotebooksClient {
	return &NotebooksClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
