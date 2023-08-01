package locations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocationsClient struct {
	Client *resourcemanager.Client
}

func NewLocationsClientWithBaseURI(sdkApi sdkEnv.Api) (*LocationsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "locations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LocationsClient: %+v", err)
	}

	return &LocationsClient{
		Client: client,
	}, nil
}
