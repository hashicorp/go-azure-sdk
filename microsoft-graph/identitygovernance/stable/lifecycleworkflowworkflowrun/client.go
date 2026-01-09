package lifecycleworkflowworkflowrun

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowRunClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowRunClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowRunClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowrun", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowRunClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowRunClient{
		Client: client,
	}, nil
}
