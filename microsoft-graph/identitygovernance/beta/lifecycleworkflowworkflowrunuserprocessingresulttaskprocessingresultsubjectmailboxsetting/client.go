package lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresultsubjectmailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultSubjectMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultSubjectMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultSubjectMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresultsubjectmailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultSubjectMailboxSettingClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultSubjectMailboxSettingClient{
		Client: client,
	}, nil
}
