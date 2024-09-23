package lifecycleworkflowworkflowversioncreatedbymailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowVersionCreatedByMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowVersionCreatedByMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowVersionCreatedByMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowversioncreatedbymailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowVersionCreatedByMailboxSettingClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowVersionCreatedByMailboxSettingClient{
		Client: client,
	}, nil
}
