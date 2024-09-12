package lifecycleworkflowrunuserprocessingresulttasksubjectmailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowRunUserProcessingResultTaskSubjectMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowRunUserProcessingResultTaskSubjectMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowRunUserProcessingResultTaskSubjectMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowrunuserprocessingresulttasksubjectmailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowRunUserProcessingResultTaskSubjectMailboxSettingClient: %+v", err)
	}

	return &LifecycleWorkflowRunUserProcessingResultTaskSubjectMailboxSettingClient{
		Client: client,
	}, nil
}
