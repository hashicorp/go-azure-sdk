package lifecycleworkflowdeleteditemworkflowcreatedby

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowDeletedItemWorkflowCreatedByClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowDeletedItemWorkflowCreatedByClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowDeletedItemWorkflowCreatedByClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowdeleteditemworkflowcreatedby", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowDeletedItemWorkflowCreatedByClient: %+v", err)
	}

	return &LifecycleWorkflowDeletedItemWorkflowCreatedByClient{
		Client: client,
	}, nil
}
