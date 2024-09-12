package lifecycleworkflowversiontask

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowVersionTaskClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowVersionTaskClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowVersionTaskClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowversiontask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowVersionTaskClient: %+v", err)
	}

	return &LifecycleWorkflowVersionTaskClient{
		Client: client,
	}, nil
}
