package lifecycleworkflowdeleteditemworkflowtaskreporttaskdefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowDeletedItemWorkflowTaskReportTaskDefinitionClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowDeletedItemWorkflowTaskReportTaskDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowDeletedItemWorkflowTaskReportTaskDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowdeleteditemworkflowtaskreporttaskdefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowDeletedItemWorkflowTaskReportTaskDefinitionClient: %+v", err)
	}

	return &LifecycleWorkflowDeletedItemWorkflowTaskReportTaskDefinitionClient{
		Client: client,
	}, nil
}
