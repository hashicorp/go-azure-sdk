package sites

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitesClient struct {
	Client *resourcemanager.Client
}

func NewSitesClientWithBaseURI(api sdkEnv.Api) (*SitesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "sites", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SitesClient: %+v", err)
	}

	return &SitesClient{
		Client: client,
	}, nil
}
