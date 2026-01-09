package lifecycleworkflowdeleteditemworkflowversionlastmodifiedby

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowdeleteditemworkflowversionlastmodifiedby", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByClient: %+v", err)
	}

	return &LifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByClient{
		Client: client,
	}, nil
}
