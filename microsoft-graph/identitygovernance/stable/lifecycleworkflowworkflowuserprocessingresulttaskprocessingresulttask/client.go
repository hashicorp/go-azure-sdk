package lifecycleworkflowworkflowuserprocessingresulttaskprocessingresulttask

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultTaskClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultTaskClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultTaskClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowuserprocessingresulttaskprocessingresulttask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultTaskClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultTaskClient{
		Client: client,
	}, nil
}
