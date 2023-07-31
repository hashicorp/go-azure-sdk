package locationcapabilities

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocationCapabilitiesClient struct {
	Client *resourcemanager.Client
}

func NewLocationCapabilitiesClientWithBaseURI(api sdkEnv.Api) (*LocationCapabilitiesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "locationcapabilities", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LocationCapabilitiesClient: %+v", err)
	}

	return &LocationCapabilitiesClient{
		Client: client,
	}, nil
}
