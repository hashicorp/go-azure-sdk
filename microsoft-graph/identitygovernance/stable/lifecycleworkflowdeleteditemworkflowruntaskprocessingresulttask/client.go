package lifecycleworkflowdeleteditemworkflowruntaskprocessingresulttask

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultTaskClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultTaskClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultTaskClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowdeleteditemworkflowruntaskprocessingresulttask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultTaskClient: %+v", err)
	}

	return &LifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultTaskClient{
		Client: client,
	}, nil
}
