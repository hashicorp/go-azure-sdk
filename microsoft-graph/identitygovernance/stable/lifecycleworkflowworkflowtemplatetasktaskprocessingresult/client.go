package lifecycleworkflowworkflowtemplatetasktaskprocessingresult

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowtemplatetasktaskprocessingresult", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultClient{
		Client: client,
	}, nil
}
