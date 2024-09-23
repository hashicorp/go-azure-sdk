package lifecycleworkflowdeleteditemworkflowtasktaskprocessingresultsubjectmailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultSubjectMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultSubjectMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultSubjectMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowdeleteditemworkflowtasktaskprocessingresultsubjectmailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultSubjectMailboxSettingClient: %+v", err)
	}

	return &LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultSubjectMailboxSettingClient{
		Client: client,
	}, nil
}
