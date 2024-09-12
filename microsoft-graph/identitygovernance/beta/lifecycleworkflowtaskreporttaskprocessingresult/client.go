package lifecycleworkflowtaskreporttaskprocessingresult

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowTaskReportTaskProcessingResultClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowTaskReportTaskProcessingResultClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowTaskReportTaskProcessingResultClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowtaskreporttaskprocessingresult", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowTaskReportTaskProcessingResultClient: %+v", err)
	}

	return &LifecycleWorkflowTaskReportTaskProcessingResultClient{
		Client: client,
	}, nil
}
