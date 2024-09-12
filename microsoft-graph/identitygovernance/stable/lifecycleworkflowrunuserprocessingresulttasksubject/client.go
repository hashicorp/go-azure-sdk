package lifecycleworkflowrunuserprocessingresulttasksubject

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowRunUserProcessingResultTaskSubjectClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowRunUserProcessingResultTaskSubjectClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowRunUserProcessingResultTaskSubjectClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowrunuserprocessingresulttasksubject", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowRunUserProcessingResultTaskSubjectClient: %+v", err)
	}

	return &LifecycleWorkflowRunUserProcessingResultTaskSubjectClient{
		Client: client,
	}, nil
}
