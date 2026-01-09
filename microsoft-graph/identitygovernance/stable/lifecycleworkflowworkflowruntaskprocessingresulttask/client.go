package lifecycleworkflowworkflowruntaskprocessingresulttask

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowRunTaskProcessingResultTaskClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowRunTaskProcessingResultTaskClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowRunTaskProcessingResultTaskClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowruntaskprocessingresulttask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowRunTaskProcessingResultTaskClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowRunTaskProcessingResultTaskClient{
		Client: client,
	}, nil
}
