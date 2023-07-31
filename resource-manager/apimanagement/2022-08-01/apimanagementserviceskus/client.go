package apimanagementserviceskus

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiManagementServiceSkusClient struct {
	Client *resourcemanager.Client
}

func NewApiManagementServiceSkusClientWithBaseURI(api sdkEnv.Api) (*ApiManagementServiceSkusClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "apimanagementserviceskus", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApiManagementServiceSkusClient: %+v", err)
	}

	return &ApiManagementServiceSkusClient{
		Client: client,
	}, nil
}
