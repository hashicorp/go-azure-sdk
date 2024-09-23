package lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresultsubjectmailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultSubjectMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultSubjectMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultSubjectMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresultsubjectmailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultSubjectMailboxSettingClient: %+v", err)
	}

	return &LifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultSubjectMailboxSettingClient{
		Client: client,
	}, nil
}
