package lifecycleworkflowrunuserprocessingresult

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowRunUserProcessingResultClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowRunUserProcessingResultClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowRunUserProcessingResultClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowrunuserprocessingresult", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowRunUserProcessingResultClient: %+v", err)
	}

	return &LifecycleWorkflowRunUserProcessingResultClient{
		Client: client,
	}, nil
}
