package workflowtriggerhistories

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkflowTriggerHistoriesClient struct {
	Client *resourcemanager.Client
}

func NewWorkflowTriggerHistoriesClientWithBaseURI(api environments.Api) (*WorkflowTriggerHistoriesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "workflowtriggerhistories", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WorkflowTriggerHistoriesClient: %+v", err)
	}

	return &WorkflowTriggerHistoriesClient{
		Client: client,
	}, nil
}
