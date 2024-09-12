package lifecycleworkflowcustomtaskextensioncreatedbyserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowCustomTaskExtensionCreatedByServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowCustomTaskExtensionCreatedByServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowCustomTaskExtensionCreatedByServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowcustomtaskextensioncreatedbyserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowCustomTaskExtensionCreatedByServiceProvisioningErrorClient: %+v", err)
	}

	return &LifecycleWorkflowCustomTaskExtensionCreatedByServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
