package lifecycleworkflowworkflowuserprocessingresult

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowUserProcessingResultClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowUserProcessingResultClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowUserProcessingResultClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowuserprocessingresult", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowUserProcessingResultClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowUserProcessingResultClient{
		Client: client,
	}, nil
}
