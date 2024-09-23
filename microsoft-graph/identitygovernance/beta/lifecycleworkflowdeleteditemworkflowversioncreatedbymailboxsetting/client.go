package lifecycleworkflowdeleteditemworkflowversioncreatedbymailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowDeletedItemWorkflowVersionCreatedByMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowDeletedItemWorkflowVersionCreatedByMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowDeletedItemWorkflowVersionCreatedByMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowdeleteditemworkflowversioncreatedbymailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowDeletedItemWorkflowVersionCreatedByMailboxSettingClient: %+v", err)
	}

	return &LifecycleWorkflowDeletedItemWorkflowVersionCreatedByMailboxSettingClient{
		Client: client,
	}, nil
}
