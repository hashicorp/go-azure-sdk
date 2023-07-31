package portalsettings

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PortalSettingsClient struct {
	Client *resourcemanager.Client
}

func NewPortalSettingsClientWithBaseURI(api environments.Api) (*PortalSettingsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "portalsettings", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PortalSettingsClient: %+v", err)
	}

	return &PortalSettingsClient{
		Client: client,
	}, nil
}
