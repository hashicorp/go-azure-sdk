package placementpolicies

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlacementPoliciesClient struct {
	Client *resourcemanager.Client
}

func NewPlacementPoliciesClientWithBaseURI(api sdkEnv.Api) (*PlacementPoliciesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "placementpolicies", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PlacementPoliciesClient: %+v", err)
	}

	return &PlacementPoliciesClient{
		Client: client,
	}, nil
}
