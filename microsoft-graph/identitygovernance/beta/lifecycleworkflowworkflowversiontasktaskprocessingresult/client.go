package lifecycleworkflowworkflowversiontasktaskprocessingresult

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowVersionTaskTaskProcessingResultClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowVersionTaskTaskProcessingResultClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowVersionTaskTaskProcessingResultClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowversiontasktaskprocessingresult", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowVersionTaskTaskProcessingResultClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowVersionTaskTaskProcessingResultClient{
		Client: client,
	}, nil
}
