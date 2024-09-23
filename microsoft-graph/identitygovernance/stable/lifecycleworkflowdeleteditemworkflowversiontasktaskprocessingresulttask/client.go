package lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresulttask

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultTaskClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultTaskClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultTaskClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresulttask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultTaskClient: %+v", err)
	}

	return &LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultTaskClient{
		Client: client,
	}, nil
}
