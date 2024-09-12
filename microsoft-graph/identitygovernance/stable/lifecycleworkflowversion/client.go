package lifecycleworkflowversion

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowVersionClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowVersionClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowVersionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowversion", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowVersionClient: %+v", err)
	}

	return &LifecycleWorkflowVersionClient{
		Client: client,
	}, nil
}
