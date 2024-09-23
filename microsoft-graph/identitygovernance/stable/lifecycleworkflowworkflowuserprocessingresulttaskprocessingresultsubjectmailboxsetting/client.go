package lifecycleworkflowworkflowuserprocessingresulttaskprocessingresultsubjectmailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultSubjectMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultSubjectMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultSubjectMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowuserprocessingresulttaskprocessingresultsubjectmailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultSubjectMailboxSettingClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultSubjectMailboxSettingClient{
		Client: client,
	}, nil
}
