package lifecycleworkflowtaskreporttask

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowTaskReportTaskClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowTaskReportTaskClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowTaskReportTaskClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowtaskreporttask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowTaskReportTaskClient: %+v", err)
	}

	return &LifecycleWorkflowTaskReportTaskClient{
		Client: client,
	}, nil
}
