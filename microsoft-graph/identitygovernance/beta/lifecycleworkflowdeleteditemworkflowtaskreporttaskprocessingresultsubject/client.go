package lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresultsubject

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresultsubject", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectClient: %+v", err)
	}

	return &LifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectClient{
		Client: client,
	}, nil
}
