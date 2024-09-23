package lifecycleworkflowworkflowtaskreporttaskprocessingresultsubjectserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowtaskreporttaskprocessingresultsubjectserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
