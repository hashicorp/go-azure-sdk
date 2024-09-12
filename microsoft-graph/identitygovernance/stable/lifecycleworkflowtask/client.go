package lifecycleworkflowtask

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowTaskClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowTaskClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowTaskClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowtask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowTaskClient: %+v", err)
	}

	return &LifecycleWorkflowTaskClient{
		Client: client,
	}, nil
}
