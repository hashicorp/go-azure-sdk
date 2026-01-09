package lifecycleworkflowworkflowversioncreatedby

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowVersionCreatedByClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowVersionCreatedByClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowVersionCreatedByClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowversioncreatedby", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowVersionCreatedByClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowVersionCreatedByClient{
		Client: client,
	}, nil
}
