package diskpoolzones

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiskPoolZonesClient struct {
	Client *resourcemanager.Client
}

func NewDiskPoolZonesClientWithBaseURI(api environments.Api) (*DiskPoolZonesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "diskpoolzones", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DiskPoolZonesClient: %+v", err)
	}

	return &DiskPoolZonesClient{
		Client: client,
	}, nil
}
