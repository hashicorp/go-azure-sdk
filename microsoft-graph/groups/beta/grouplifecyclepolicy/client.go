package grouplifecyclepolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupLifecyclePolicyClient struct {
	Client *msgraph.Client
}

func NewGroupLifecyclePolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupLifecyclePolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouplifecyclepolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupLifecyclePolicyClient: %+v", err)
	}

	return &GroupLifecyclePolicyClient{
		Client: client,
	}, nil
}
