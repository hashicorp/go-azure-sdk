package lifecycleworkflowdeleteditemworkflowexecutionscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowDeletedItemWorkflowExecutionScopeClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowDeletedItemWorkflowExecutionScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowDeletedItemWorkflowExecutionScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowdeleteditemworkflowexecutionscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowDeletedItemWorkflowExecutionScopeClient: %+v", err)
	}

	return &LifecycleWorkflowDeletedItemWorkflowExecutionScopeClient{
		Client: client,
	}, nil
}
