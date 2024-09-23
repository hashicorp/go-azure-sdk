package lifecycleworkflowdeleteditemworkflowtask

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowDeletedItemWorkflowTaskClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowDeletedItemWorkflowTaskClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowDeletedItemWorkflowTaskClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowdeleteditemworkflowtask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowDeletedItemWorkflowTaskClient: %+v", err)
	}

	return &LifecycleWorkflowDeletedItemWorkflowTaskClient{
		Client: client,
	}, nil
}
