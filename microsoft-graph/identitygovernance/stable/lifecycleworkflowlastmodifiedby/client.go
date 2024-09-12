package lifecycleworkflowlastmodifiedby

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowLastModifiedByClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowLastModifiedByClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowLastModifiedByClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowlastmodifiedby", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowLastModifiedByClient: %+v", err)
	}

	return &LifecycleWorkflowLastModifiedByClient{
		Client: client,
	}, nil
}
