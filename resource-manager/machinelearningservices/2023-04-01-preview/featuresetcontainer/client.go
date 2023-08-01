package featuresetcontainer

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeaturesetContainerClient struct {
	Client *resourcemanager.Client
}

func NewFeaturesetContainerClientWithBaseURI(sdkApi sdkEnv.Api) (*FeaturesetContainerClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "featuresetcontainer", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating FeaturesetContainerClient: %+v", err)
	}

	return &FeaturesetContainerClient{
		Client: client,
	}, nil
}
