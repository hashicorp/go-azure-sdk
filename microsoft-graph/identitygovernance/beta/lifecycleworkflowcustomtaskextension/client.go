package lifecycleworkflowcustomtaskextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowCustomTaskExtensionClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowCustomTaskExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowCustomTaskExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowcustomtaskextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowCustomTaskExtensionClient: %+v", err)
	}

	return &LifecycleWorkflowCustomTaskExtensionClient{
		Client: client,
	}, nil
}
