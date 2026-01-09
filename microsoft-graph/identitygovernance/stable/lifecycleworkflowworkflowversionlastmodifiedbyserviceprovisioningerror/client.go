package lifecycleworkflowworkflowversionlastmodifiedbyserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowVersionLastModifiedByServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowVersionLastModifiedByServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowVersionLastModifiedByServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowversionlastmodifiedbyserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowVersionLastModifiedByServiceProvisioningErrorClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowVersionLastModifiedByServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
