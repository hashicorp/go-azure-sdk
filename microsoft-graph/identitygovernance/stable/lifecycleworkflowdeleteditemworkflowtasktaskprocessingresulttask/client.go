package lifecycleworkflowdeleteditemworkflowtasktaskprocessingresulttask

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultTaskClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultTaskClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultTaskClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowdeleteditemworkflowtasktaskprocessingresulttask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultTaskClient: %+v", err)
	}

	return &LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultTaskClient{
		Client: client,
	}, nil
}
