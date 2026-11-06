package availableenvironmentmodes

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AvailableEnvironmentModesClient struct {
	Client *resourcemanager.Client
}

func NewAvailableEnvironmentModesClientWithBaseURI(sdkApi sdkEnv.Api) (*AvailableEnvironmentModesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "availableenvironmentmodes", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AvailableEnvironmentModesClient: %+v", err)
	}

	return &AvailableEnvironmentModesClient{
		Client: client,
	}, nil
}
