package pipelineruns

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PipelineRunsClient struct {
	Client *dataplane.Client
}

func NewPipelineRunsClientUnconfigured() (*PipelineRunsClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "pipelineruns", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PipelineRunsClient: %+v", err)
	}

	return &PipelineRunsClient{
		Client: client,
	}, nil
}

func (c *PipelineRunsClient) PipelineRunsClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewPipelineRunsClientWithBaseURI(endpoint string) (*PipelineRunsClient, error) {
	client, err := dataplane.NewClient(endpoint, "pipelineruns", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PipelineRunsClient: %+v", err)
	}

	return &PipelineRunsClient{
		Client: client,
	}, nil
}

func (c *PipelineRunsClient) Clone(endpoint string) *PipelineRunsClient {
	return &PipelineRunsClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
