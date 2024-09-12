package lifecycleworkflowcreatedby

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowCreatedByClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowCreatedByClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowCreatedByClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowcreatedby", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowCreatedByClient: %+v", err)
	}

	return &LifecycleWorkflowCreatedByClient{
		Client: client,
	}, nil
}
