package workflows

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkflowsClient struct {
	Client *resourcemanager.Client
}

func NewWorkflowsClientWithBaseURI(api environments.Api) (*WorkflowsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "workflows", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WorkflowsClient: %+v", err)
	}

	return &WorkflowsClient{
		Client: client,
	}, nil
}
