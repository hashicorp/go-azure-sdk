package lifecycleworkflowworkflowrunuserprocessingresultsubject

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowRunUserProcessingResultSubjectClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowRunUserProcessingResultSubjectClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowRunUserProcessingResultSubjectClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowrunuserprocessingresultsubject", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowRunUserProcessingResultSubjectClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowRunUserProcessingResultSubjectClient{
		Client: client,
	}, nil
}
