package jobschedules

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobSchedulesClient struct {
	Client *dataplane.Client
}

func NewJobSchedulesClientUnconfigured() (*JobSchedulesClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "jobschedules", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JobSchedulesClient: %+v", err)
	}

	return &JobSchedulesClient{
		Client: client,
	}, nil
}

func (c *JobSchedulesClient) JobSchedulesClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewJobSchedulesClientWithBaseURI(endpoint string) (*JobSchedulesClient, error) {
	client, err := dataplane.NewClient(endpoint, "jobschedules", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JobSchedulesClient: %+v", err)
	}

	return &JobSchedulesClient{
		Client: client,
	}, nil
}

func (c *JobSchedulesClient) Clone(endpoint string) *JobSchedulesClient {
	return &JobSchedulesClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
