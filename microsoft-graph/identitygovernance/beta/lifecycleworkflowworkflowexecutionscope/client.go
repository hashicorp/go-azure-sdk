package lifecycleworkflowworkflowexecutionscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowExecutionScopeClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowExecutionScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowExecutionScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowexecutionscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowExecutionScopeClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowExecutionScopeClient{
		Client: client,
	}, nil
}
