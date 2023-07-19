package zone

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ZoneClient struct {
	Client *resourcemanager.Client
}

func NewZoneClientWithBaseURI(api environments.Api) (*ZoneClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "zone", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ZoneClient: %+v", err)
	}

	return &ZoneClient{
		Client: client,
	}, nil
}
