package lifecycleworkflowworkflowversion

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowVersionClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowVersionClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowVersionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowversion", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowVersionClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowVersionClient{
		Client: client,
	}, nil
}
