package lifecycleworkflowworkflowtemplatetask

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowWorkflowTemplateTaskClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowWorkflowTemplateTaskClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowWorkflowTemplateTaskClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowworkflowtemplatetask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowWorkflowTemplateTaskClient: %+v", err)
	}

	return &LifecycleWorkflowWorkflowTemplateTaskClient{
		Client: client,
	}, nil
}
