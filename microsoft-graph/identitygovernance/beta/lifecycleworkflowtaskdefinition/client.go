package lifecycleworkflowtaskdefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowTaskDefinitionClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowTaskDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowTaskDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowtaskdefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowTaskDefinitionClient: %+v", err)
	}

	return &LifecycleWorkflowTaskDefinitionClient{
		Client: client,
	}, nil
}
