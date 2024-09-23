package lifecycleworkflowdeleteditemworkflowversiontask

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowDeletedItemWorkflowVersionTaskClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowDeletedItemWorkflowVersionTaskClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowDeletedItemWorkflowVersionTaskClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowdeleteditemworkflowversiontask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowDeletedItemWorkflowVersionTaskClient: %+v", err)
	}

	return &LifecycleWorkflowDeletedItemWorkflowVersionTaskClient{
		Client: client,
	}, nil
}
