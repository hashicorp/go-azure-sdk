package lifecycleworkflowdeleteditemworkflowruntaskprocessingresultsubjectserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultSubjectServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultSubjectServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultSubjectServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowdeleteditemworkflowruntaskprocessingresultsubjectserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultSubjectServiceProvisioningErrorClient: %+v", err)
	}

	return &LifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultSubjectServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
