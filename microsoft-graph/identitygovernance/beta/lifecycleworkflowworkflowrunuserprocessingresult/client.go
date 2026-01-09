package lifecycleworkflowworkflowrunuserprocessingresult

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowRunUserProcessingResultClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowRunUserProcessingResultClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowRunUserProcessingResultClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowrunuserprocessingresult", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowRunUserProcessingResultClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowRunUserProcessingResultClient{
		Client: client,
	}, nil
}
