package lifecycleworkflowworkflowtasktaskprocessingresultsubjectserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowTaskTaskProcessingResultSubjectServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowTaskTaskProcessingResultSubjectServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowTaskTaskProcessingResultSubjectServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowtasktaskprocessingresultsubjectserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowTaskTaskProcessingResultSubjectServiceProvisioningErrorClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowTaskTaskProcessingResultSubjectServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
