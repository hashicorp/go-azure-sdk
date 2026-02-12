package tasks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TasksClient struct {
	Client *dataplane.Client
}

func NewTasksClientUnconfigured() (*TasksClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "tasks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TasksClient: %+v", err)
	}

	return &TasksClient{
		Client: client,
	}, nil
}

func (c *TasksClient) TasksClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewTasksClientWithBaseURI(endpoint string) (*TasksClient, error) {
	client, err := dataplane.NewClient(endpoint, "tasks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TasksClient: %+v", err)
	}

	return &TasksClient{
		Client: client,
	}, nil
}

func (c *TasksClient) Clone(endpoint string) *TasksClient {
	return &TasksClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
