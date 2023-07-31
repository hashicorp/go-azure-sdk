package portalconfig

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PortalConfigClient struct {
	Client *resourcemanager.Client
}

func NewPortalConfigClientWithBaseURI(api sdkEnv.Api) (*PortalConfigClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "portalconfig", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PortalConfigClient: %+v", err)
	}

	return &PortalConfigClient{
		Client: client,
	}, nil
}
