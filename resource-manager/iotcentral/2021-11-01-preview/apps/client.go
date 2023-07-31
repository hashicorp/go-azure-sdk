package apps

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppsClient struct {
	Client *resourcemanager.Client
}

func NewAppsClientWithBaseURI(api sdkEnv.Api) (*AppsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "apps", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AppsClient: %+v", err)
	}

	return &AppsClient{
		Client: client,
	}, nil
}
