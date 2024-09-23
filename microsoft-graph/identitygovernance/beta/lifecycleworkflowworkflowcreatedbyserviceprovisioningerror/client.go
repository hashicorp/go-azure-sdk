package lifecycleworkflowworkflowcreatedbyserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowCreatedByServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowCreatedByServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowCreatedByServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowcreatedbyserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowCreatedByServiceProvisioningErrorClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowCreatedByServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
