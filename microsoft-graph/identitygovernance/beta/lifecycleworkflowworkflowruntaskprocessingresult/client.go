package lifecycleworkflowworkflowruntaskprocessingresult

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowRunTaskProcessingResultClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowRunTaskProcessingResultClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowRunTaskProcessingResultClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowruntaskprocessingresult", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowRunTaskProcessingResultClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowRunTaskProcessingResultClient{
		Client: client,
	}, nil
}
