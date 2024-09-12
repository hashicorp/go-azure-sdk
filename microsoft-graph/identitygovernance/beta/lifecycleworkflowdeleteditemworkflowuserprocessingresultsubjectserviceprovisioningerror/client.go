package lifecycleworkflowdeleteditemworkflowuserprocessingresultsubjectserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowdeleteditemworkflowuserprocessingresultsubjectserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectServiceProvisioningErrorClient: %+v", err)
	}

	return &LifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
