package featuresetversion

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeaturesetVersionClient struct {
	Client *resourcemanager.Client
}

func NewFeaturesetVersionClientWithBaseURI(api sdkEnv.Api) (*FeaturesetVersionClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "featuresetversion", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating FeaturesetVersionClient: %+v", err)
	}

	return &FeaturesetVersionClient{
		Client: client,
	}, nil
}
