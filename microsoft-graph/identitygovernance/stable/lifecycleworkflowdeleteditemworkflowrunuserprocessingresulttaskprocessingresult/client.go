package lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresult

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresult", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultClient: %+v", err)
	}

	return &LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultClient{
		Client: client,
	}, nil
}
