package lifecycleworkflowrunuserprocessingresultsubject

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowRunUserProcessingResultSubjectClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowRunUserProcessingResultSubjectClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowRunUserProcessingResultSubjectClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowrunuserprocessingresultsubject", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowRunUserProcessingResultSubjectClient: %+v", err)
	}

	return &LifecycleWorkflowRunUserProcessingResultSubjectClient{
		Client: client,
	}, nil
}
