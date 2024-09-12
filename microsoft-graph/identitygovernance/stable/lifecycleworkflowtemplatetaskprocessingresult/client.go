package lifecycleworkflowtemplatetaskprocessingresult

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowTemplateTaskProcessingResultClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowTemplateTaskProcessingResultClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowTemplateTaskProcessingResultClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowtemplatetaskprocessingresult", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowTemplateTaskProcessingResultClient: %+v", err)
	}

	return &LifecycleWorkflowTemplateTaskProcessingResultClient{
		Client: client,
	}, nil
}
