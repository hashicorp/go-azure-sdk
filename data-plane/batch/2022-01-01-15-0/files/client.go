package files

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FilesClient struct {
	Client *dataplane.Client
}

func NewFilesClientUnconfigured() (*FilesClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "files", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating FilesClient: %+v", err)
	}

	return &FilesClient{
		Client: client,
	}, nil
}

func (c *FilesClient) FilesClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewFilesClientWithBaseURI(endpoint string) (*FilesClient, error) {
	client, err := dataplane.NewClient(endpoint, "files", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating FilesClient: %+v", err)
	}

	return &FilesClient{
		Client: client,
	}, nil
}

func (c *FilesClient) Clone(endpoint string) *FilesClient {
	return &FilesClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
