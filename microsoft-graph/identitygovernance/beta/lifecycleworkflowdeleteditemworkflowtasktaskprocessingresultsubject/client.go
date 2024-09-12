package lifecycleworkflowdeleteditemworkflowtasktaskprocessingresultsubject

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultSubjectClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultSubjectClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultSubjectClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowdeleteditemworkflowtasktaskprocessingresultsubject", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultSubjectClient: %+v", err)
	}

	return &LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultSubjectClient{
		Client: client,
	}, nil
}
