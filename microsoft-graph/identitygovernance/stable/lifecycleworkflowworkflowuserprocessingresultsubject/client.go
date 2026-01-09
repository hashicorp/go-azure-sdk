package lifecycleworkflowworkflowuserprocessingresultsubject

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowUserProcessingResultSubjectClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowUserProcessingResultSubjectClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowUserProcessingResultSubjectClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowuserprocessingresultsubject", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowUserProcessingResultSubjectClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowUserProcessingResultSubjectClient{
		Client: client,
	}, nil
}
