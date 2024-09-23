package lifecycleworkflowcustomtaskextensionlastmodifiedbymailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowCustomTaskExtensionLastModifiedByMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowCustomTaskExtensionLastModifiedByMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowCustomTaskExtensionLastModifiedByMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowcustomtaskextensionlastmodifiedbymailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowCustomTaskExtensionLastModifiedByMailboxSettingClient: %+v", err)
	}

	return &LifecycleWorkflowCustomTaskExtensionLastModifiedByMailboxSettingClient{
		Client: client,
	}, nil
}
