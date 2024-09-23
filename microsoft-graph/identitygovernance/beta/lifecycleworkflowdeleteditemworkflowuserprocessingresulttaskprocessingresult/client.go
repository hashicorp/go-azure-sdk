package lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresult

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresult", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultClient: %+v", err)
	}

	return &LifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultClient{
		Client: client,
	}, nil
}
