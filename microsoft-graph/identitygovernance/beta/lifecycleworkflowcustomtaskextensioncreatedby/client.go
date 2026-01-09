package lifecycleworkflowcustomtaskextensioncreatedby

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowCustomTaskExtensionCreatedByClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowCustomTaskExtensionCreatedByClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowCustomTaskExtensionCreatedByClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowcustomtaskextensioncreatedby", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowCustomTaskExtensionCreatedByClient: %+v", err)
	}

	return &LifecycleWorkflowCustomTaskExtensionCreatedByClient{
		Client: client,
	}, nil
}
