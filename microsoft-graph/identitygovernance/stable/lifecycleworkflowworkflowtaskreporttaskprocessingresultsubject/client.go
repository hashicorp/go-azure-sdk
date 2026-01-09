package lifecycleworkflowworkflowtaskreporttaskprocessingresultsubject

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowtaskreporttaskprocessingresultsubject", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectClient{
		Client: client,
	}, nil
}
