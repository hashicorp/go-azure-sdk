package feature

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeatureClient struct {
	Client *resourcemanager.Client
}

func NewFeatureClientWithBaseURI(api environments.Api) (*FeatureClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "feature", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating FeatureClient: %+v", err)
	}

	return &FeatureClient{
		Client: client,
	}, nil
}
