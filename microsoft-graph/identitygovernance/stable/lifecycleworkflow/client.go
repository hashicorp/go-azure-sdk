package lifecycleworkflow

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflow", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowClient: %+v", err)
	}

	return &LifecycleWorkflowClient{
		Client: client,
	}, nil
}
