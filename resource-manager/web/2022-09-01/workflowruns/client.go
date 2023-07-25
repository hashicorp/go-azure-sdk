package workflowruns

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkflowRunsClient struct {
	Client *resourcemanager.Client
}

func NewWorkflowRunsClientWithBaseURI(api environments.Api) (*WorkflowRunsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "workflowruns", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WorkflowRunsClient: %+v", err)
	}

	return &WorkflowRunsClient{
		Client: client,
	}, nil
}
