package migrates

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigratesClient struct {
	Client *resourcemanager.Client
}

func NewMigratesClientWithBaseURI(api environments.Api) (*MigratesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "migrates", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MigratesClient: %+v", err)
	}

	return &MigratesClient{
		Client: client,
	}, nil
}
