package lifecycleworkflowcustomtaskextensioncreatedbymailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowCustomTaskExtensionCreatedByMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowCustomTaskExtensionCreatedByMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowCustomTaskExtensionCreatedByMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowcustomtaskextensioncreatedbymailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowCustomTaskExtensionCreatedByMailboxSettingClient: %+v", err)
	}

	return &LifecycleWorkflowCustomTaskExtensionCreatedByMailboxSettingClient{
		Client: client,
	}, nil
}
