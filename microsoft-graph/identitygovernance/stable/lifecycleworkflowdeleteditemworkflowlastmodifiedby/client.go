package lifecycleworkflowdeleteditemworkflowlastmodifiedby

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowDeletedItemWorkflowLastModifiedByClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowDeletedItemWorkflowLastModifiedByClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowDeletedItemWorkflowLastModifiedByClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowdeleteditemworkflowlastmodifiedby", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowDeletedItemWorkflowLastModifiedByClient: %+v", err)
	}

	return &LifecycleWorkflowDeletedItemWorkflowLastModifiedByClient{
		Client: client,
	}, nil
}
