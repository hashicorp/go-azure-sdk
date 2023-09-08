package groupgrouplifecyclepolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupGroupLifecyclePolicyClient struct {
	Client *msgraph.Client
}

func NewGroupGroupLifecyclePolicyClientWithBaseURI(api sdkEnv.Api) (*GroupGroupLifecyclePolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupgrouplifecyclepolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupGroupLifecyclePolicyClient: %+v", err)
	}

	return &GroupGroupLifecyclePolicyClient{
		Client: client,
	}, nil
}
