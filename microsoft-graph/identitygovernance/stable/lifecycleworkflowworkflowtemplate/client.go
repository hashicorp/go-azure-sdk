package lifecycleworkflowworkflowtemplate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowTemplateClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowTemplateClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowTemplateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowtemplate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowTemplateClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowTemplateClient{
		Client: client,
	}, nil
}
