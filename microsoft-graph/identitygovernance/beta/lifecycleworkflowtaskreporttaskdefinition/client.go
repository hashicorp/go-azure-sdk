package lifecycleworkflowtaskreporttaskdefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowTaskReportTaskDefinitionClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowTaskReportTaskDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowTaskReportTaskDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowtaskreporttaskdefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowTaskReportTaskDefinitionClient: %+v", err)
	}

	return &LifecycleWorkflowTaskReportTaskDefinitionClient{
		Client: client,
	}, nil
}
