package lifecycleworkflowworkflowcreatedby

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowCreatedByClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowCreatedByClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowCreatedByClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowcreatedby", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowCreatedByClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowCreatedByClient{
		Client: client,
	}, nil
}
