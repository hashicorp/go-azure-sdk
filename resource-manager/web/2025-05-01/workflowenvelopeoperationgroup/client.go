package workflowenvelopeoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkflowEnvelopeOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewWorkflowEnvelopeOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*WorkflowEnvelopeOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "workflowenvelopeoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WorkflowEnvelopeOperationGroupClient: %+v", err)
	}

	return &WorkflowEnvelopeOperationGroupClient{
		Client: client,
	}, nil
}
