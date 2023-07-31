package simpolicies

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SIMPoliciesClient struct {
	Client *resourcemanager.Client
}

func NewSIMPoliciesClientWithBaseURI(api sdkEnv.Api) (*SIMPoliciesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "simpolicies", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SIMPoliciesClient: %+v", err)
	}

	return &SIMPoliciesClient{
		Client: client,
	}, nil
}
