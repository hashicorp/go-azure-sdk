package workflowrunactionrepetitiondefinitions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkflowRunActionRepetitionDefinitionsClient struct {
	Client *resourcemanager.Client
}

func NewWorkflowRunActionRepetitionDefinitionsClientWithBaseURI(sdkApi sdkEnv.Api) (*WorkflowRunActionRepetitionDefinitionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "workflowrunactionrepetitiondefinitions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WorkflowRunActionRepetitionDefinitionsClient: %+v", err)
	}

	return &WorkflowRunActionRepetitionDefinitionsClient{
		Client: client,
	}, nil
}
