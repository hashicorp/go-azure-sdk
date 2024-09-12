package lifecycleworkflowversioncreatedby

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowVersionCreatedByClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowVersionCreatedByClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowVersionCreatedByClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowversioncreatedby", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowVersionCreatedByClient: %+v", err)
	}

	return &LifecycleWorkflowVersionCreatedByClient{
		Client: client,
	}, nil
}
