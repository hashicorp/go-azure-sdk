package lifecycleworkflowworkflowruntaskprocessingresultsubject

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowRunTaskProcessingResultSubjectClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowRunTaskProcessingResultSubjectClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowRunTaskProcessingResultSubjectClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowruntaskprocessingresultsubject", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowRunTaskProcessingResultSubjectClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowRunTaskProcessingResultSubjectClient{
		Client: client,
	}, nil
}
