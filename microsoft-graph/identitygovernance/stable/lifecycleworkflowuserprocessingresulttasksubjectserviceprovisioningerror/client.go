package lifecycleworkflowuserprocessingresulttasksubjectserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowUserProcessingResultTaskSubjectServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowUserProcessingResultTaskSubjectServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowUserProcessingResultTaskSubjectServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowuserprocessingresulttasksubjectserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowUserProcessingResultTaskSubjectServiceProvisioningErrorClient: %+v", err)
	}

	return &LifecycleWorkflowUserProcessingResultTaskSubjectServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
