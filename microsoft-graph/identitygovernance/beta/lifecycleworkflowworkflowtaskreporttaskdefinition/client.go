package lifecycleworkflowworkflowtaskreporttaskdefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowTaskReportTaskDefinitionClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowTaskReportTaskDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowTaskReportTaskDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowtaskreporttaskdefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowTaskReportTaskDefinitionClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowTaskReportTaskDefinitionClient{
		Client: client,
	}, nil
}
