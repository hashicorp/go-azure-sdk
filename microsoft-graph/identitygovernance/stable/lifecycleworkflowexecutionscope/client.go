package lifecycleworkflowexecutionscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowExecutionScopeClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowExecutionScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowExecutionScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowexecutionscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowExecutionScopeClient: %+v", err)
	}

	return &LifecycleWorkflowExecutionScopeClient{
		Client: client,
	}, nil
}
