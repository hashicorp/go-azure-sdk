package lifecycleworkflowversioncreatedbymailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowVersionCreatedByMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowVersionCreatedByMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowVersionCreatedByMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowversioncreatedbymailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowVersionCreatedByMailboxSettingClient: %+v", err)
	}

	return &LifecycleWorkflowVersionCreatedByMailboxSettingClient{
		Client: client,
	}, nil
}
