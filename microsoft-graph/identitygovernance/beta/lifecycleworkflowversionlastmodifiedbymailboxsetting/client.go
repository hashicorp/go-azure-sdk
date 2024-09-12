package lifecycleworkflowversionlastmodifiedbymailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowVersionLastModifiedByMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowVersionLastModifiedByMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowVersionLastModifiedByMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowversionlastmodifiedbymailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowVersionLastModifiedByMailboxSettingClient: %+v", err)
	}

	return &LifecycleWorkflowVersionLastModifiedByMailboxSettingClient{
		Client: client,
	}, nil
}
