package lifecycleworkflowversionlastmodifiedby

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleWorkflowVersionLastModifiedByClient struct {
	Client *msgraph.Client
}

func NewLifecycleWorkflowVersionLastModifiedByClientWithBaseURI(sdkApi sdkEnv.Api) (*LifecycleWorkflowVersionLastModifiedByClient, error) {
	client, err := msgraph.NewClient(sdkApi, "lifecycleworkflowversionlastmodifiedby", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LifecycleWorkflowVersionLastModifiedByClient: %+v", err)
	}

	return &LifecycleWorkflowVersionLastModifiedByClient{
		Client: client,
	}, nil
}
