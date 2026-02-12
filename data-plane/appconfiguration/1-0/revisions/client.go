package revisions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RevisionsClient struct {
	Client *dataplane.Client
}

func NewRevisionsClientUnconfigured() (*RevisionsClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "revisions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RevisionsClient: %+v", err)
	}

	return &RevisionsClient{
		Client: client,
	}, nil
}

func (c *RevisionsClient) RevisionsClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewRevisionsClientWithBaseURI(endpoint string) (*RevisionsClient, error) {
	client, err := dataplane.NewClient(endpoint, "revisions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RevisionsClient: %+v", err)
	}

	return &RevisionsClient{
		Client: client,
	}, nil
}

func (c *RevisionsClient) Clone(endpoint string) *RevisionsClient {
	return &RevisionsClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
