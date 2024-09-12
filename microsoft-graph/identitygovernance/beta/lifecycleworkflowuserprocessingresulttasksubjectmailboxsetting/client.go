package lifecycleworkflowuserprocessingresulttasksubjectmailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowUserProcessingResultTaskSubjectMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowUserProcessingResultTaskSubjectMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowUserProcessingResultTaskSubjectMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowuserprocessingresulttasksubjectmailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowUserProcessingResultTaskSubjectMailboxSettingClient: %+v", err)
	}

	return &LifecycleWorkflowUserProcessingResultTaskSubjectMailboxSettingClient{
		Client: client,
	}, nil
}
