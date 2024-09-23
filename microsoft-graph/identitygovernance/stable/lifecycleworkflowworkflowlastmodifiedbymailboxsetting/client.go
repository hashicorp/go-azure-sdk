package lifecycleworkflowworkflowlastmodifiedbymailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowLastModifiedByMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowLastModifiedByMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowLastModifiedByMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowlastmodifiedbymailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowLastModifiedByMailboxSettingClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowLastModifiedByMailboxSettingClient{
		Client: client,
	}, nil
}
