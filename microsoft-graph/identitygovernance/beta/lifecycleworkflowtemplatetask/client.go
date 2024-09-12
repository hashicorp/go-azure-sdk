package lifecycleworkflowtemplatetask

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowTemplateTaskClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowTemplateTaskClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowTemplateTaskClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowtemplatetask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowTemplateTaskClient: %+v", err)
	}

	return &LifecycleWorkflowTemplateTaskClient{
		Client: client,
	}, nil
}
