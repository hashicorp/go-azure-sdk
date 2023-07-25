package policies

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PoliciesClient struct {
	Client *resourcemanager.Client
}

func NewPoliciesClientWithBaseURI(api environments.Api) (*PoliciesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "policies", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PoliciesClient: %+v", err)
	}

	return &PoliciesClient{
		Client: client,
	}, nil
}
