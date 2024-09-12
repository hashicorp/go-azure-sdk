package lifecycleworkflowdeleteditemworkflowrun

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowDeletedItemWorkflowRunClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowDeletedItemWorkflowRunClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowDeletedItemWorkflowRunClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowdeleteditemworkflowrun", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowDeletedItemWorkflowRunClient: %+v", err)
	}

	return &LifecycleWorkflowDeletedItemWorkflowRunClient{
		Client: client,
	}, nil
}
