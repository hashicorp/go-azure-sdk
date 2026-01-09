package lifecycleworkflowsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowSettingClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowSettingClient: %+v", err)
	}

	return &LifecycleWorkflowSettingClient{
		Client: client,
	}, nil
}
